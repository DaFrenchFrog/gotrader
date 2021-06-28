package model

import "time"

//Response :
type Response struct {
	Success bool `json:"success"`
}

//CoinListResponse :
type CoinListResponse struct {
	Response
	Result []MarketData `json:"result"`
}

//MarketDataResponse :
type MarketDataResponse struct {
	Response
	Result MarketData
}

//CoinHistoryResponse :
type CoinHistoryResponse struct {
	Response
	Result []Candle
}

//Candle :
type Candle struct {
	Close      float32   `json:"close"`
	High       float32   `json:"high"`
	Low        float32   `json:"low"`
	Open       float32   `json:"open"`
	ClockTime  float32   `json:"time"`
	StartTime  time.Time `json:"startTime"`
	Volume     float32   `json:"volume"`
	Resolution string
}

//MarketData :
type MarketData struct {
	Name           string  `json:"name"`
	BaseCurrency   string  `json:"baseCurrency"`
	QuoteCurrency  string  `json:"quoteCurrency"`
	Type           string  `json:"type"`
	Underlying     string  `json:"underlying"`
	Enabled        bool    `json:"enabled"`
	Ask            float32 `json:"ask"`
	Bid            float32 `json:"bid"`
	Last           float32 `json:"last"`
	PostOnly       bool    `json:"postOnly"`
	PriceIncrement float32 `json:"priceIncrement"`
	SizeIncrement  float32 `json:"sizeIncrement"`
	Restricted     bool    `json:"restricted"`
	History        []Candle
}

// Layout format for date logging
const DateLayoutLog = "02 Jan 2006 15h04"

// Color get color
func Color(c string) string {

	switch c {
	case "red":
		return string("\033[31m")
	case "green":
		return string("\033[32m")
	case "yellow":
		return string("\033[33m")
	case "blue":
		return string("\033[34m")
	case "purple":
		return string("\033[35m")
	case "cyan":
		return string("\033[36m")
	case "white":
		return string("\033[37m")
	default:
		return string("\033[0m")
	}
}
