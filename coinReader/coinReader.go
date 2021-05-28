package coinReader

import (
	"fmt"
	"strings"

	"github.com/elRomano/gotrader/apiCaller"
)

type response struct {
	Success bool `json:"success"`
}

type coinListResponse struct {
	response
	Result []coinData `json:"result"`
}

type coinDataResponse struct {
	response //object composition
	result   coinData
}

type coinData struct {
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
}

const baseURL = "https://ftx.com/api/markets"

// CoinReader is
type CoinReader struct {
}

// New is
func New() CoinReader {
	return CoinReader{}
}

func (c CoinReader) ListCoin(coin string) {
	resp := &coinDataResponse{}
	apiCaller.Call(baseURL+"/"+coin, resp)
	fmt.Printf("Succes:%v", resp.Success)
	fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", resp.result.Name, resp.result.PriceIncrement)
}

// ListMarkets is...
func (c CoinReader) ListMarkets(coin string) {

	resp := &coinListResponse{}
	apiCaller.Call(baseURL, resp)

	for _, r := range resp.Result {
		if strings.Contains(r.Name, coin) {
			fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", r.Name, r.PriceIncrement)
		}
	}
}
