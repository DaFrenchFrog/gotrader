package strategy

import (
	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/model"
)

//NormalStrategy :
type NormalStrategy struct {
	cumulativeGreen int
	cumulativeRed   int
}

func (s NormalStrategy) init() {
	s.cumulativeGreen = 0
	s.cumulativeRed = 0
}

func (s *NormalStrategy) apply(w *account.Wallet, ticker model.CoinHistoryDataTicker) {
	// cfmt.Println("-> ", &s.cumulativeGreen)
	if ticker.Open > ticker.Close {
		s.cumulativeGreen = 0
		if w.Trading {
			// cfmt.Println(cfmt.Red, "SELL : ", v.Close)
			w.Trading = false
			w.WalletAmount += w.Opening - ticker.Close
		}
		// cfmt.Println(cfmt.Red, "Red...")
		//fmt.Println(model.Color("red"), "Red...", model.Color(""))
	} else {
		// cfmt.Println(cfmt.Green, "Green...")
		s.cumulativeGreen++
		if s.cumulativeGreen >= 3 && !w.Trading {
			// cfmt.Println(cfmt.Green, "BUY : ", v.Close)
			w.Trading = true
			w.Opening = ticker.Close
			w.TradeAmount++
		}
	}
}
