package strategy

import (
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
}

//New :
func New(s strategy, w account.Wallet) StrategyRunner {
	return StrategyRunner{
		wallet:   w,
		strategy: s,
	}
}

//Backtest :
func (s StrategyRunner) Backtest(market string) {
	reader := markets.NewReader(market)
	reader.Load()

	s.strategy.init()

	cfmt.Println(cfmt.Purple, "Starting backtest : ", s.marketData.Name+" -> testing ", len(s.marketData.History), " entries", cfmt.Neutral)

	for _, v := range s.marketData.History {
		s.strategy.apply(&s.wallet, v)
	}
	cfmt.Println(cfmt.Blue, "BACKTEST FINISHED : WINNINGS : ", s.wallet.WalletAmount, "$ // TRADES : ", s.wallet.TradeAmount)
}

//Live :
func (s *StrategyRunner) Live(market string) {
	s.reader = markets.NewReader(market)
	s.reader.Load()

	second := time.Tick(time.Second)
	newCandle := s.reader.NewCandleChannel()

	for {
		select {
		case <-second:
			s.reader.GetLatestCandle()
			// cfmt.Println(cfmt.Blue, curMarket.Market.Last)
		case <-newCandle:
			s.strategy.apply(&s.wallet, newCandle)
			// cfmt.Println(cfmt.Blue, curMarket.Market.Last)
		}
	}
}
