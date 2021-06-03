package coinReader

import (
	"fmt"

	"github.com/elRomano/gotrader/apiCaller"
	"github.com/elRomano/gotrader/model"
)

const baseURL = "https://ftx.com/api"

// CoinReader is
type CoinReader struct {
}

// New is
func New() CoinReader {
	return CoinReader{}
}

// GetHistory is
func (c CoinReader) GetHistory(coin string) error {
	resp := &model.CoinHistoryResponse{}
	succeed, err := apiCaller.Call(baseURL+"/markets/"+coin+"/candles?resolution=60", resp)
	if err != nil {
		return err
	}
	if !succeed {
		fmt.Println("No error but no success either...")
		return nil
	}

	for _, r := range resp.Result {
		fmt.Printf("{\tname: %v\tbaseCurrency: %v\tclockTime: %v}\n", r.Close, r.Volume, r.ClockTime)
	}
	fmt.Println(len(resp.Result))
	return nil
}

//ListCoin is
func (c CoinReader) ListCoin(coin string) error {
	resp := &model.CoinDataResponse{}
	succeed, err := apiCaller.Call(baseURL+"/markets/"+coin, resp)

	if err != nil {
		return err
	}
	if !succeed {
		fmt.Println("No error but no success either...")
		return nil
	}

	fmt.Printf("name: %v\tbaseCurrency: %v\n", resp.Result.Name, resp.Result.BaseCurrency)

	return nil
}

// ListMarkets is...
func (c CoinReader) ListMarkets() error {
	resp := &model.CoinListResponse{}
	succeed, err := apiCaller.Call(baseURL+"/markets", resp)

	if err != nil {
		return err
	}
	if !succeed {
		fmt.Println("No error but no success either...")
		return nil
	}

	for _, r := range resp.Result {
		fmt.Printf("{\tname: %v\tbaseCurrency: %v}\n", r.Name, r.BaseCurrency)
	}
	return nil
}
