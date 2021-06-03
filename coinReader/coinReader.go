package coinReader

import (
	"fmt"

	"github.com/elRomano/gotrader/apiCaller"
	"github.com/elRomano/gotrader/model"
)

const baseURL = "https://ftx.com/api/markets"

// CoinReader is
type CoinReader struct {
}

// New is
func New() CoinReader {
	return CoinReader{}
}

func (c CoinReader) getHistory(coin string) error {
	return nil
}

// Les fonction retournent des erreur plutot que de crasher le prog. C'est la responsabilité du main de crasher pas d'un package que tu appel
func (c CoinReader) ListCoin(coin string) error {
	resp := &model.CoinDataResponse{}
	succeed, err := apiCaller.Call(baseURL+"/"+coin, resp)

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
	succeed, err := apiCaller.Call(baseURL, resp)

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
