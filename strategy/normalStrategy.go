package strategy

import (
	"math"
	"strings"

	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

//NormalStrategy :
type NormalStrategy struct {
	cumulativeGreen int
	cumulativeRed   int
	WinLoseTrades   map[bool]int
	position        Position
	logged          int
}

//NewNormalStrategy :
func NewNormalStrategy() *NormalStrategy {
	return &NormalStrategy{
		cumulativeGreen: 0,
		cumulativeRed:   0,
		WinLoseTrades:   map[bool]int{},
		logged:          0,
	}
}

func (s *NormalStrategy) init() {
	s.cumulativeGreen = 0
	s.cumulativeRed = 0
	s.WinLoseTrades = map[bool]int{}
	s.position = Position{}
}

func (s *NormalStrategy) apply(w *account.Wallet, candle model.Candle, market string) {
	// showTicker(candle)

	if candle.Open > candle.Close {
		// cfmt.Println(cfmt.Red, "ATR14 : ", candle.ATR14)

		s.cumulativeGreen = 0
		if s.position != (Position{}) {
			isWinning := false
			w.Balance[market], isWinning = s.position.Close(candle.Close)
			s.WinLoseTrades[isWinning]++
		}
	} else {
		// cfmt.Println(cfmt.Green, "Green...")
		s.cumulativeGreen++
		if s.cumulativeGreen >= 3 && s.position == (Position{}) {
			// cfmt.Println(cfmt.Green, "BUY : ", v.Close)
			s.position = NewPosition("long", candle.Close, (w.Balance[market]*.1)/candle.Close)
			w.Balance[market] = 0
		}
	}
}

//WinningTrades :
func (s *NormalStrategy) WinningTrades() int {
	return s.WinLoseTrades[true]
}

//LosingTrades :
func (s *NormalStrategy) LosingTrades() int {
	return s.WinLoseTrades[false]
}

//logCandle
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
