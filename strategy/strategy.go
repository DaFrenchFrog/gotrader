package strategy

import (
	"fmt"

	coinreader "github.com/elRomano/gotrader/coinReader"
	"github.com/elRomano/gotrader/model"
)

type Strategy struct {
}

//New :
func New() Strategy {
	return Strategy{}
}

//Backtest :
func (s Strategy) Backtest(history coinreader.CoinReader) {
	fmt.Printf(model.Color("purple"), "starting backtest...", model.Color(""))

}
