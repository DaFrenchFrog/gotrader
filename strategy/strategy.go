package strategy

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/elRomano/gotrader/store"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/markets"
	"github.com/elRomano/gotrader/model"
)

type strategy interface {
	init(walletAmount float32)
	apply(ticker model.Candle, market string)
	WinningTrades() int
	LosingTrades() int
	GetBalance() float32
	GetClosedPositions() []Position
}

//StrategyRunner :
type StrategyRunner struct {
	strategies   map[string]strategy
	reader       markets.MarketReader
	walletAmount float32
	// marketData []model.MarketData
	// store store.HistoryStore
}

//New :
func New(marketList []string, store store.HistoryStore, w float32) StrategyRunner {
	return StrategyRunner{
		strategies:   map[string]strategy{},
		reader:       markets.NewReader(marketList, store),
		walletAmount: w,
	}
}

//LaunchBacktest :
func (s *StrategyRunner) LaunchBacktest(term string) {
	err := s.reader.Load(term)
	if err != nil {
		log.Fatalf("Error loading data, %v", err)
	}
	fmt.Println("starting range")
	for _, mData := range s.reader.Data {
		fmt.Println("starting init", mData.Name)
		s.strategies[mData.Name] = NewSupertrendStrategy()
		s.strategies[mData.Name].init(s.walletAmount)
		candleAmount := len(s.reader.Data[mData.Name].History)
		cfmt.Println(cfmt.Blue, "Starting backtest : ", mData.Name+" -> testing ", candleAmount, " entries from ", mData.History["1m"][0].StartTime.Format(model.DateLayoutLog), cfmt.Neutral)

		for _, v := range s.reader.Data[mData.Name].History["1m"] {
			s.strategies[mData.Name].apply(v, mData.Name)
		}
		dailyAverage := candleAmount / 60 / 24
		winnings := s.strategies[mData.Name].GetBalance() - s.walletAmount
		for pos, h := range s.strategies[mData.Name].GetClosedPositions() {
			cfmt.Println(cfmt.Cyan, "n°", pos, "\t", h.PositionType, "\t O:", h.OpeningPrice, "\t C:", h.ClosingPrice, "\t W:", h.Winnings)
		}

		cfmt.
			Println(cfmt.Purple, "BACKTEST FINISHED : ", cfmt.Neutral, int(winnings/float32(dailyAverage)), "$/day \t", winnings, "$ total \t win/lose trades : ", s.strategies[mData.Name].WinningTrades(), "/", s.strategies[mData.Name].LosingTrades())
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
func logCandle(c model.Candle) {
	if c.Open > c.Close {
		cfmt.Print(cfmt.Red, c.StartTime.Format(model.DateLayoutLog))
	} else {
		cfmt.Print(cfmt.Green, c.StartTime.Format(model.DateLayoutLog))
	}
	cfmt.Println(cfmt.Neutral, " O=", c.Open, " H=", c.High, " L=", c.Low, " C=", c.Close, " VOL=", c.Volume, " \t ATR7=", c.ATR7, " SMA200=", c.SMA200, "\t ST : bLO=", c.STrend.BasicLowerBand, " bUP=", c.STrend.BasicUpperBand, " V=", c.STrend.Color, c.STrend.Value, "$")
}

//showCandle
func showCandle(candle model.Candle) {
	drawingLength := float32(10)
	var thinLowRatio float64
	var thickRatio float64
	var thinHighRatio float64
	var drawing string
	amplitude := candle.High - candle.Low
	drawing = ""

	if candle.Open < candle.Close {
		//GREEN
		thinLowRatio = math.Round(float64((candle.Open - candle.Low) / amplitude * drawingLength))
		thickRatio = math.Round(float64((candle.Close - candle.Open) / amplitude * drawingLength))
		thinHighRatio = math.Round(float64((candle.High - candle.Close) / amplitude * drawingLength))
		// cfmt.Println(cfmt.Red, thinLowRatio, " + ", thickRatio, " + ", thinHighRatio, " = ", thinLowRatio+thickRatio+thinHighRatio)
	}
	drawing += "["
	drawing += strings.Repeat("-", int(thinLowRatio))
	drawing += strings.Repeat("=", int(thickRatio))
	drawing += strings.Repeat("-", int(thinHighRatio))
	drawing += "]"
	cfmt.Println(cfmt.Green, (candle.Open - candle.Low), "/", amplitude, "*", drawingLength, " = \t", thinLowRatio, "\t", thinHighRatio, "\t", thickRatio, "\t", drawing)
}
