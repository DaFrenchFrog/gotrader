package gfx

import (
	"image/color"
	"math"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
	"github.com/fogleman/gg"
)

//DrawChart :
func DrawChart(candles []model.Candle) error {
	const W = 6000
	const H = 2000
	const SPACING = 4
	dc := gg.NewContext(W, H)
	dc.SetColor(color.Black)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()
	min, max := getCandleRange(candles, W/SPACING)
	height := float64(max - min)
	dc.SetLineWidth(1)
	candlesAmount := 0
	for pos, c := range candles {
		//DRAW CANDLES
		dc.SetLineWidth(1)
		if c.Open < c.Close {
			dc.SetRGB(0, .8, 0)
		} else {
			dc.SetRGB(.8, 0, 0)
		}
		dc.DrawLine(float64(pos*SPACING), H-float64(math.Round(float64(c.High-min)/height*H)), float64(pos*SPACING), H-float64(c.Low-min)/height*H)
		dc.Stroke()
		dc.SetLineWidth(3)
		dc.DrawLine(float64(pos*SPACING), H-float64(math.Round(float64(c.Open-min)/height*H)), float64(pos*SPACING), H-float64(c.Close-min)/height*H)
		dc.Stroke()

		if pos > 200 {
			//DRAW SMA200
			dc.SetLineWidth(2)
			dc.SetRGB(.5, .5, 1)
			dc.DrawLine(float64((pos-1)*SPACING), H-float64(math.Round(float64(candles[pos-1].SMA200-min)/height*H)), float64(pos*SPACING), H-float64(math.Round(float64(c.SMA200-min)/height*H)))
			dc.Stroke()

			//DRAW SUPERTREND
			// dc.SetColor(color.Color)
			if c.STrend.Color == cfmt.Green {
				dc.SetRGB(0, .8, 0)
			} else {
				dc.SetRGB(.8, 0, 0)
			}
			dc.DrawLine(float64((pos-1)*SPACING), H-float64(math.Round(float64(candles[pos-1].STrend.Value-min)/height*H)), float64(pos*SPACING), H-float64(math.Round(float64(c.STrend.Value-min)/height*H)))
			// dc.DrawCircle(float64(pos*SPACING), H-float64(math.Round(float64(c.STrend.Value-min)/height*H)), SPACING/2)
			dc.Stroke()
		}

		candlesAmount = pos
		if pos*SPACING > W {
			break
		}

	}
	//DRAW TEXT
	dc.SetColor(color.White)
	s := candles[0].StartTime.Format(model.DateLayoutLog) + " -> " + candles[candlesAmount].StartTime.Format(model.DateLayoutLog)
	dc.DrawString(s, 10, 10)

	gg.SaveJPG("out.jpg", dc.Image(), 80)
	return nil
}

func getCandleRange(c []model.Candle, r int) (float32, float32) {
	min := float32(999999999999)
	max := float32(0)
	for i := range c {
		if c[i].Low < min {
			min = c[i].Low
		}
		if c[i].High > max {
			max = c[i].High
		}
	}
	return min, max
}
