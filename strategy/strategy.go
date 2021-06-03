package strategy

import "fmt"

type Strategy struct {
}

func New() Strategy {
	return Strategy{}
}

func (s Strategy) Backtest() {
	fmt.Println("hey nice test !")
}
