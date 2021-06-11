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
}

func NewNormalStrategy() *NormalStrategy {
	return &NormalStrategy{
		cumulativeGreen: 0,
		cumulativeRed:   0,
	}
}

func (s *NormalStrategy) init() {
	s.cumulativeGreen = 0
	s.cumulativeRed = 0
}

func (s *NormalStrategy) apply(w *account.Wallet, ticker model.Candle) {
	// cfmt.Println("-> ", &s.cumulativeGreen)
	// showTicker(ticker)
	if ticker.Open > ticker.Close {
		// cfmt.Println(cfmt.Red, "SELL : ", ticker.Close)
		s.cumulativeGreen = 0
		if w.Trading {

			w.Trading = false
			w.WalletAmount += w.Opening - ticker.Close
		}
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

func showTicker(ticker model.Candle) {
	drawingLength := float32(10)
	var thinLowRatio float64
	var thickRatio float64
	var thinHighRatio float64
	var drawing string
	amplitude := ticker.High - ticker.Low
	drawing = ""

	if ticker.Open < ticker.Close {
		//GREEN
		thinLowRatio = math.Round(float64((ticker.Open - ticker.Low) / amplitude * drawingLength))
		thickRatio = math.Round(float64((ticker.Close - ticker.Open) / amplitude * drawingLength))
		thinHighRatio = math.Round(float64((ticker.High - ticker.Close) / amplitude * drawingLength))
		// cfmt.Println(cfmt.Red, thinLowRatio, " + ", thickRatio, " + ", thinHighRatio, " = ", thinLowRatio+thickRatio+thinHighRatio)
	}
	drawing += "["
	drawing += strings.Repeat("-", int(thinLowRatio))
	drawing += strings.Repeat("=", int(thickRatio))
	drawing += strings.Repeat("-", int(thinHighRatio))
	drawing += "]"
	cfmt.Println(cfmt.Green, (ticker.Open - ticker.Low), "/", amplitude, "*", drawingLength, " = \t", thinLowRatio, "\t", thinHighRatio, "\t", thickRatio, "\t", drawing)
}
