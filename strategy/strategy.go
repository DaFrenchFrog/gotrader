package strategy

import (
	"fmt"
	"github.com/elRomano/gotrader/cfmt"

	"github.com/elRomano/gotrader/model"
)

type Strategy struct {
}

//New :
func New() Strategy {
	return Strategy{}
}

//Backtest :
func (s Strategy) Backtest(market model.CoinData) {
	fmt.Println(model.Color("purple"), "Starting backtest : ", market.Name+" -> testing", len(market.History), " entries", model.Color(""))

	for _, v := range market.History {
		if v.Open > v.Close {
			cfmt.Println(cfmt.Red, "Red...")
			//fmt.Println(model.Color("red"), "Red...", model.Color(""))
		} else {
			cfmt.Println(cfmt.Green, "Green...")
			//fmt.Println(model.Color("green"), "Green...", model.Color(""))
		}

	}

}
