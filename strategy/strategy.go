package strategy

import (
	"fmt"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

type Strategy struct {
	walletAmount    float32
	cumulativeGreen int
	cumulativeRed   int
	trading         bool
	tradeAmount     int
	opening         float32
}

//New :
func New() Strategy {
	return Strategy{}
}

//Backtest :
func (s Strategy) Backtest(market model.CoinData) {
	s.cumulativeGreen = 0
	s.tradeAmount = 0
	fmt.Println(model.Color("purple"), "Starting backtest : ", market.Name+" -> testing", len(market.History), " entries", model.Color(""))

	for _, v := range market.History {
		if v.Open > v.Close {
			s.cumulativeGreen = 0
			if s.trading {
				// cfmt.Println(cfmt.Red, "SELL : ", v.Close)
				s.trading = false
				s.walletAmount += s.opening - v.Close
			}
			// cfmt.Println(cfmt.Red, "Red...")
			//fmt.Println(model.Color("red"), "Red...", model.Color(""))
		} else {
			// cfmt.Println(cfmt.Green, "Green...")
			s.cumulativeGreen++
			if s.cumulativeGreen >= 3 && !s.trading {
				// cfmt.Println(cfmt.Green, "BUY : ", v.Close)
				s.trading = true
				s.opening = v.Close
				s.tradeAmount++
			}
			//fmt.Println(model.Color("green"), "Green...", model.Color(""))
		}

	}
	cfmt.Println(cfmt.Blue, "BACKTEST FINISHED : WINNINGS : ", s.walletAmount, "$ // TRADES : ", s.tradeAmount)

}
