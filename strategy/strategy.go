package strategy

import (
	"log"
	"time"

	"github.com/elRomano/gotrader/store"

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
	wallet   account.Wallet
	strategy strategy
	reader   markets.MarketReader
	// marketData []model.MarketData
	// store store.HistoryStore
}

//New :
func New(marketList []string, s strategy, w account.Wallet, store store.HistoryStore) StrategyRunner {
	return StrategyRunner{
		wallet:   w,
		strategy: s,
		reader:   markets.NewReader(marketList, store),
		// store:    store,
	}
}

//LaunchBacktest :
func (s *StrategyRunner) LaunchBacktest() {

	err := s.reader.Load()
	if err != nil {
		log.Fatalf("Error loading data, %v", err)
	}
	s.strategy.init()
	for _, mData := range s.reader.Data {
		cfmt.Println(cfmt.Blue, "Starting backtest : ", mData.Name+" -> testing ", len(s.reader.Data[mData.Name].History), " entries", cfmt.Neutral)

		for _, v := range s.reader.Data[mData.Name].History {
			s.strategy.apply(&s.wallet, v)
		}
		cfmt.Println(cfmt.Purple, "BACKTEST FINISHED : WINNINGS : ", s.wallet.WalletAmount, "$ // TRADES : ", s.wallet.TradeAmount)

	}

	// dataFromDb, err := s.store.GetAll()
	// if err != nil {
	// 	cfmt.Printf(cfmt.Red, "Shit, %v", err)
	// }
	// for _, candle := range dataFromDb {
	// 	fmt.Println(candle)
	// }
}

//Live :
func (s *StrategyRunner) Live(market []string) {

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
			// s.reader.GetLatestCandle()
		}
	}

}
