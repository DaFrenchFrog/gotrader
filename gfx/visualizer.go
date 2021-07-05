package gfx

import (
	"image/color"
	"math"

	"github.com/elRomano/gotrader/model"
	"github.com/fogleman/gg"
)

//DrawChart :
func DrawChart(candles []model.Candle) {
	const W = 6000
	const H = 2000
	const SPACING = 4
	dc := gg.NewContext(W, H)
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()
	min, max := getCandleRange(candles, W/SPACING)
	height := float64(max - min)
	dc.SetLineWidth(1)
	for pos, c := range candles {
		//DRAW CANDLES
		dc.SetLineWidth(1)
		// fmt.Println((c.Open < c.Close), c.Open, c.Close)
		if c.Open < c.Close {
			dc.SetRGB(0, .8, 0)
		} else {
			dc.SetRGB(.8, 0, 0)
		}
		dc.DrawLine(float64(pos*SPACING), float64(math.Round(float64(c.High-min)/height*H)), float64(pos*SPACING), float64(c.Low-min)/height*H)
		dc.Stroke()
		dc.SetLineWidth(3)
		dc.DrawLine(float64(pos*SPACING), float64(math.Round(float64(c.Open-min)/height*H)), float64(pos*SPACING), float64(c.Close-min)/height*H)
		dc.Stroke()
		//DRAW SMA200
		dc.SetRGB(.5, .5, 1)
		dc.DrawCircle(float64(pos*SPACING), float64(math.Round(float64(c.SMA200-min)/height*H)), SPACING/2)
		dc.Fill()
		//DRAW SUPERTREND
		// dc.SetColor(color.Color)
		dc.DrawCircle(float64(pos*SPACING), float64(math.Round(float64(c.STrend.Value-min)/height*H)), SPACING/2)
		dc.Fill()
		if pos*SPACING > W {
			break
		}
	}
	gg.SaveJPG("out.jpg", dc.Image(), 80)
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
