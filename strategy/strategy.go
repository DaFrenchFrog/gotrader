package strategy

import (
	"fmt"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

type wallet struct {
	walletAmount    float32
	cumulativeGreen int
	cumulativeRed   int
	trading         bool
	tradeAmount     int
	opening         float32
}

type strategy interface {
	apply(w *wallet, ticker model.CoinHistoryDataTicker)
}

type StrategyRunner struct {
	wallet
	strategy strategy
}

//New :
func New(s strategy) StrategyRunner {
	return StrategyRunner{
		wallet:   wallet{},
		strategy: s,
	}
}

//Backtest :
func (s StrategyRunner) Backtest(market model.CoinData) {
	s.cumulativeGreen = 0
	s.tradeAmount = 0
	fmt.Println(model.Color("purple"), "Starting backtest : ", market.Name+" -> testing", len(market.History), " entries", model.Color(""))

	for _, v := range market.History {
		s.strategy.apply(&s.wallet, v)
	}
	cfmt.Println(cfmt.Blue, "BACKTEST FINISHED : WINNINGS : ", s.walletAmount, "$ // TRADES : ", s.tradeAmount)
}

type NormalStrategy struct {
}

func (NormalStrategy) apply(w *wallet, ticker model.CoinHistoryDataTicker) {
	if ticker.Open > ticker.Close {
		w.cumulativeGreen = 0
		if w.trading {
			// cfmt.Println(cfmt.Red, "SELL : ", v.Close)
			w.trading = false
			w.walletAmount += w.opening - ticker.Close
		}
		// cfmt.Println(cfmt.Red, "Red...")
		//fmt.Println(model.Color("red"), "Red...", model.Color(""))
	} else {
		// cfmt.Println(cfmt.Green, "Green...")
		w.cumulativeGreen++
		if w.cumulativeGreen >= 3 && !w.trading {
			// cfmt.Println(cfmt.Green, "BUY : ", v.Close)
			w.trading = true
			w.opening = ticker.Close
			w.tradeAmount++
		}
		//fmt.Println(model.Color("green"), "Green...", model.Color(""))
	}
}

type CrazyStrategy struct {
}

func (CrazyStrategy) apply(w *wallet, ticker model.CoinHistoryDataTicker) {
	if ticker.Open < ticker.Close {
		w.cumulativeGreen = 0
		if w.trading {
			w.trading = false
			w.walletAmount += w.opening - ticker.Close
		}
	} else {
		w.cumulativeGreen++
		if w.cumulativeGreen >= 3 && !w.trading {
			w.trading = true
			w.opening = ticker.Close
			w.tradeAmount++
		}
	}
}
