package markets

import (
	"fmt"
	"strings"

	"github.com/elRomano/gotrader/ftx"
)

//MarketLister :
type MarketLister struct {
	client ftx.Client
}

// NewLister :
func NewLister() MarketLister {
	return MarketLister{
		client: ftx.Client{},
	}
}

//ListMarkets :
func (m MarketLister) ListMarkets() error {
	resp, err := m.client.ListMarkets()

	if err != nil {
		return err
	}
	if !resp.Success {
		fmt.Println("No error but no success either...")
		return nil
	}

	for _, r := range resp.Result {
		// fmt.Printf("{\tname: %v\tbaseCurrency: %v}\n", r.Name, r.BaseCurrency)
		if strings.Contains(r.Name, "BULL/USDT") {
			fmt.Print(r.Name, " ")
		}
	}
	return nil
}
