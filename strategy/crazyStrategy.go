package strategy

import (
	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/model"
)

// CrazyStrategy :
type CrazyStrategy struct {
}

func (CrazyStrategy) init() {

}

func (CrazyStrategy) apply(w *account.Wallet, ticker model.CoinHistoryDataTicker) {
	if ticker.Open < ticker.Close {
		// BUY
	} else {
		// SELL
	}
}
