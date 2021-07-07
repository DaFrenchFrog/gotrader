package strategy

import (
	"github.com/elRomano/gotrader/model"
)

//NormalStrategy :
type NormalStrategy struct {
	cumulativeGreen int
	cumulativeRed   int
	WinLoseTrades   map[bool]int
	position        Position
	logged          int
	Balance         float32
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

func (s *NormalStrategy) init(w float32) {
	s.cumulativeGreen = 0
	s.cumulativeRed = 0
	s.WinLoseTrades = map[bool]int{}
	s.position = Position{}
	s.Balance = w
}

func (s *NormalStrategy) apply(candle model.Candle, market string) {
	// showTicker(candle)

	if candle.Open > candle.Close {
		// cfmt.Println(cfmt.Red, "ATR14 : ", candle.ATR14)

		s.cumulativeGreen = 0
		if s.position != (Position{}) {
			isWinning := false
			s.Balance, isWinning = s.position.Close(candle.Close, candle.StartTime)
			s.WinLoseTrades[isWinning]++
		}
	} else {
		// cfmt.Println(cfmt.Green, "Green...")
		s.cumulativeGreen++
		if s.cumulativeGreen >= 3 && s.position == (Position{}) {
			// cfmt.Println(cfmt.Green, "BUY : ", v.Close)
			s.position = NewPosition("long", candle.Close, (s.Balance*.1)/candle.Close, candle.StartTime)
			s.Balance = 0
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
