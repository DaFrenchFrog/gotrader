package strategy

import (
	"fmt"
	"github.com/elRomano/gotrader/store"
	"time"

	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/markets"
	"github.com/elRomano/gotrader/model"
)

type strategy interface {
	init()
	apply(w *account.Wallet, ticker model.Candle)
}

//StrategyRunner :
type StrategyRunner struct {
	wallet     account.Wallet
	strategy   strategy
	reader     markets.MarketReader
	marketData model.MarketData
	store      store.HistoryStore
}

//New :
func New(s strategy, w account.Wallet, store store.HistoryStore) StrategyRunner {
	return StrategyRunner{
		wallet:   w,
		strategy: s,
		store:    store,
	}
}

//Backtest :
func (s *StrategyRunner) Backtest(market string) {
	reader := markets.NewReader(market, s.store)
	reader.Load()
	s.marketData = reader.Data
	s.strategy.init()

	cfmt.Println(cfmt.Purple, "Starting backtest : ", s.marketData.Name+" -> testing ", len(s.marketData.History), " entries", cfmt.Neutral)

	for _, v := range s.marketData.History {
		s.strategy.apply(&s.wallet, v)
	}
	cfmt.Println(cfmt.Blue, "BACKTEST FINISHED : WINNINGS : ", s.wallet.WalletAmount, "$ // TRADES : ", s.wallet.TradeAmount)

	dataFromDb, err := s.store.GetAll()
	if err != nil {
		cfmt.Printf(cfmt.Red, "Shit, %v", err)
	}
	for _, candle := range dataFromDb {
		fmt.Println(candle)
	}
}

//Live :
func (s *StrategyRunner) Live(market string) {
	s.reader = markets.NewReader(market, s.store)
	s.reader.Load()

	second := time.Tick(time.Second)
	newCandle := s.reader.GetNewCandleChannel()

	go func() {
		for candle := range newCandle {
			cfmt.Printf(cfmt.Blue, "New candle!! high: %v low:%v", candle.High, candle.Low)
			s.strategy.apply(&s.wallet, candle)
		}
	}()

	for {
		for range second {
			s.reader.GetLatestCandle()
		}
	}

}
