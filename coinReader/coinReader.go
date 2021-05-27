package coinReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elRomano/gotrader/apiCaller"
)

type response struct {
	Success bool     `json:"success"`
	Result  []result `json:"result"`
}

type result struct {
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
	httpResp, err := http.Get(baseURL + "/" + coin)
	if err != nil {
		log.Fatal(err)
	}
}

// ListMarkets is...
func (c CoinReader) ListMarkets(coin string) {

	apiCaller.Call()
	apiCaller.Call()

	for _, r := range resp.Result {
		if strings.Contains(r.Name, coin) {
			fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", r.Name, r.PriceIncrement)
		}
	}
}

func (c CoinReader) Read() {
	httpResp, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer httpResp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(httpResp.Body)
	bodyString := string(bodyBytes)
	fmt.Printf(bodyString)
	var resp = response{}
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp success: %v \n", resp.Success)

	for _, r := range resp.Result {
		if strings.Contains(r.Name, "USDT") {
			fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", r.Name, r.PriceIncrement)
		}
	}
}
