package coinReader

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elRomano/gotrader/apiCaller"
)

type coinlListType struct {
	Success bool     `json:"success"`
	Result  []result `json:"result"`
}

type coinDataType struct {
	Name           string  `json:"name"`
	BaseCurrency   string  `json:baseCurrency`
	QuoteCurrency  string  `json:quoteCurrency`
	Type           string  `json:type`
	Underlying     string  `json:underlying`
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
	apiCaller.Call(baseURL+"/"+coin, reflect.TypeOf((*coinDataType)(nil)))
}

// ListMarkets is...
func (c CoinReader) ListMarkets(coin string) {

	apiCaller.Call(baseURL, coinlListType)

	for _, r := range resp.Result {
		if strings.Contains(r.Name, coin) {
			fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", r.Name, r.PriceIncrement)
		}
	}
}
