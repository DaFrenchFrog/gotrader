package coinreader

import (
	"fmt"

	"github.com/elRomano/gotrader/apiCaller"
	"github.com/elRomano/gotrader/model"
)

const baseURL = "https://ftx.com/api"

// CoinReader :
type CoinReader struct {
	Market model.CoinData
}

// New :
func New() CoinReader {
	return CoinReader{}
}

//GetCoinData :
func (c CoinReader) GetCoinData(coin string) error {
	c.listCoin(coin)
	c.GetCoinHistory(coin)
	return nil
}

//GetCoinHistory :
func (c CoinReader) GetCoinHistory(coin string) error {
	resp := &model.CoinHistoryResponse{}
	succeed, err := apiCaller.Call(baseURL+"/markets/"+coin+"/candles?resolution=60", resp)
	if err != nil {
		return err
	}
	if !succeed {
		fmt.Println("No error but no success either...")
		return nil
	}
	resultLength := len(resp.Result)
	fmt.Println(model.Color("green"), "Tickers loaded : ", model.Color(""), resultLength, coin, " entries from", resp.Result[0].StartTime.String(), " to ", resp.Result[resultLength-1].StartTime.String())

	c.Market.History = resp.Result
	// for _, r := range resp.Result {
	// 	fmt.Printf("{\tname: %v\tbaseCurrency: %v\tclockTime: %v}\n", r.Close, r.Volume, r.StartTime.String())
	// }
	return nil
}

//ListCoin :
func (c CoinReader) listCoin(coin string) error {
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

//ListMarkets :
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
