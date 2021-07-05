package markets

import (
	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/util"

	"github.com/elRomano/gotrader/model"
)

func getSMA(period int, candles []model.Candle, n int) float32 {
	if n < period {
		return 0
	}
	total := float32(0)
	for i := n - period; i < n; i++ {
		total += candles[i].Close
	}
	// fmt.Println(total / period)
	return total / float32(period)
}

func getATR(period int, c []model.Candle, n int) float32 {
	if n < period {
		return 0
	}
	total := float32(0)
	for i := n - period; i < n; i++ {
		total += util.Ftoabs(c[i].High - c[i].Low)
		// fmt.Println(i, period)
	}
	return total / float32(period)
}

func getSupertrend(multiplier int, c []model.Candle, n int) model.SuperTrend {
	s := model.SuperTrend{}

	s.BasicUpperBand = (c[n].High+c[n].Low)/2 + float32(multiplier)*c[n].ATR7
	s.BasicLowerBand = (c[n].High+c[n].Low)/2 - float32(multiplier)*c[n].ATR7

	if n < 200 {
		s.FinalUpperBand = 0
		s.FinalLowerBand = 0
	} else {
		// fmt.Println(n, "basiclo= ", s.BasicLowerBand, " \tbasicup=", s.BasicUpperBand, " \tfinalLO-1=", c[n-1].STrend.FinalUpperBand, " \tfinalUP-1=", c[n-1].STrend.FinalUpperBand)
		if s.BasicUpperBand < c[n-1].STrend.FinalUpperBand || c[n-1].Close > c[n-1].STrend.FinalUpperBand {
			s.FinalUpperBand = s.BasicUpperBand
		} else {
			s.FinalUpperBand = c[n-1].STrend.FinalUpperBand
		}
		if s.BasicLowerBand > c[n-1].STrend.FinalLowerBand || c[n-1].Close < c[n-1].STrend.FinalLowerBand {
			s.FinalLowerBand = s.BasicLowerBand
		} else {
			s.FinalLowerBand = c[n-1].STrend.FinalLowerBand
		}

		if c[n-1].STrend.Value == c[n-1].STrend.FinalUpperBand && c[n].Close <= s.FinalUpperBand {
			s.Value = s.FinalUpperBand
			s.Color = cfmt.Red
		} else if c[n-1].STrend.Value == c[n-1].STrend.FinalUpperBand && c[n].Close >= s.FinalUpperBand {
			s.Value = s.FinalLowerBand
			s.Color = cfmt.Green
		} else if c[n-1].STrend.Value == c[n-1].STrend.FinalLowerBand && c[n].Close >= s.FinalLowerBand {
			s.Value = s.FinalLowerBand
			s.Color = cfmt.Green
		} else if c[n-1].STrend.Value == c[n-1].STrend.FinalLowerBand && c[n].Close <= s.FinalLowerBand {
			s.Value = s.FinalUpperBand
			s.Color = cfmt.Red
		} else {
			s.Value = 0
		}
	}
	return s
}
