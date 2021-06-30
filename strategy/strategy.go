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
	apply(w *account.Wallet, ticker model.Candle, market string)
	WinningTrades() int
	LosingTrades() int
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
	for _, mData := range s.reader.Data {
		s.strategy.init()
		candleAmount := len(s.reader.Data[mData.Name].History)
		cfmt.Println(cfmt.Blue, "Starting backtest : ", mData.Name+" -> testing ", candleAmount, " entries from ", mData.History[0].StartTime.Format(model.DateLayoutLog), cfmt.Neutral)

		for _, v := range s.reader.Data[mData.Name].History {
			s.strategy.apply(&s.wallet, v, mData.Name)
		}
		dailyAverage := candleAmount / 60 / 24
		cfmt.
			Println(cfmt.Purple, "BACKTEST FINISHED : ", cfmt.Neutral, int(s.wallet.Balance[mData.Name]/float32(dailyAverage)), "$/day \t", s.wallet.Balance[mData.Name], "$ total \t win/lose trades : ", s.strategy.WinningTrades(), "/", s.strategy.LosingTrades())
	}
}

//Live :
func (s *StrategyRunner) Live(market []string) {

	second := time.Tick(time.Second)
	newCandle := s.reader.GetNewCandleChannel()

	go func() {
		for candle := range newCandle {
			cfmt.Printf(cfmt.Blue, "New candle!! high: %v low:%v", candle.High, candle.Low)
			// s.strategy.apply(&s.wallet, candle,)
		}
	}()

	for {
		for range second {
			// s.reader.GetLatestCandle()
		}
	}
}
