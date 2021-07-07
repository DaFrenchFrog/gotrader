package strategy

import (
	"fmt"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

//SupertrendStrategy :
type SupertrendStrategy struct {
	prevTrend       model.SuperTrend
	WinLoseTrades   map[bool]int
	position        Position
	closedPositions []Position
	logged          int
	Balance         float32
}

//NewSupertrendStrategy :
func NewSupertrendStrategy() *SupertrendStrategy {
	return &SupertrendStrategy{
		prevTrend:       model.SuperTrend{},
		logged:          0,
		position:        Position{},
		closedPositions: []Position{},
		WinLoseTrades:   map[bool]int{},
	}
}

func (s *SupertrendStrategy) init(wallet float32) {
	fmt.Println("SET WALLET")
	s.Balance = wallet
}

func (s *SupertrendStrategy) apply(candle model.Candle, market string) {
	// showTicker(candle)
	if s.prevTrend.Value > 0 {
		if s.prevTrend.Color != candle.STrend.Color {
			if s.position.open {
				cash, isWinning := s.position.Close(candle.Close, candle.StartTime)
				s.closedPositions = append(s.closedPositions, s.position)
				s.Balance += cash
				s.WinLoseTrades[isWinning]++
			}
			if candle.STrend.Color == cfmt.Green {
				// cfmt.Println(cfmt.Cyan, "BUY : ", candle.Close)
				s.position = NewPosition("long", candle.Close, (s.Balance*.1)/candle.Close, candle.StartTime)
				s.Balance -= s.Balance * .1
			} else {
				s.position = NewPosition("short", candle.Close, (s.Balance*.1)/candle.Close, candle.StartTime)
				s.Balance -= s.Balance * .1
			}
		}
	}
	s.prevTrend = candle.STrend
}

//WinningTrades :
func (s *SupertrendStrategy) WinningTrades() int {
	return s.WinLoseTrades[true]
}

//LosingTrades :
func (s *SupertrendStrategy) LosingTrades() int {
	return s.WinLoseTrades[false]
}

//GetBalance :
func (s *SupertrendStrategy) GetBalance() float32 {
	return s.Balance
}

//GetClosedPositions :
func (s *SupertrendStrategy) GetClosedPositions() []Position {
	return s.closedPositions
}
