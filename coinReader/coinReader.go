package coinreader

import (
	"fmt"
	"github.com/elRomano/gotrader/ftx"
	"time"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

// CoinReader :
type CoinReader struct {
	Market model.CoinData
	client ftx.Client
}

// New :
func New() CoinReader {
	return CoinReader{
		client: ftx.Client{},
	}
}

//GetCoinData :
func (c *CoinReader) GetCoinData(coin string) error {
	err := c.loadMarket(coin)
	if err != nil {
		return err
	}
	return c.GetCoinHistory(coin)
}

//GetCoinHistory :
func (c *CoinReader) GetCoinHistory(coin string) error {
	startDateToLoad := time.Now().AddDate(0, -1, 0)
	y, m, d := startDateToLoad.Date()
	cfmt.Println(cfmt.Purple, "Starting history loading : ", cfmt.Neutral, "starting date ", d, m, y)
	fmt.Println("Starting time : ", startDateToLoad)
	return c.getFramedCoinHistory(coin, startDateToLoad.Unix())
}

//getFramedCoinHistory :
func (c *CoinReader) getFramedCoinHistory(coin string, dateStart int64) error {
	timeframe := 60
	endDate := dateStart + int64(1500*timeframe)

	resp, err := c.client.GetCoinHistory(coin, 60, dateStart, endDate)
	if err != nil {
		return err
	}

	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error...")
		return nil
	}
	resultLength := len(resp.Result)
	cfmt.Println(cfmt.Green, "Tickers loaded : ", cfmt.Neutral, resultLength, " ", coin, " entries from \t", resp.Result[0].StartTime.Format("2 Jan 2006 15:04"), " to \t", resp.Result[resultLength-1].StartTime.Format("2 Jan 2006 15:04"))

	c.appendHistory(resp.Result)
	if endDate < time.Now().Unix() {
		c.getFramedCoinHistory(coin, endDate)
	}
	return nil
}
func (c *CoinReader) appendHistory(history []model.CoinHistoryDataTicker) {
	c.Market.History = append(c.Market.History, history...)
}

func (c *CoinReader) loadMarket(coin string) error {
	resp, err := c.client.ListCoin(coin)

	if err != nil {
		return err
	}
	if !resp.Success {
		fmt.Println("No error but no success either...")
		return nil
	}

	fmt.Println(model.Color("green"), "Market loaded: ", model.Color(""), " name: ", resp.Result.Name, "\tbaseCurrency: ", resp.Result.BaseCurrency)
	c.Market = resp.Result
	return nil
}

//ListMarkets :
func (c CoinReader) ListMarkets() error {
	resp, err := c.client.ListMarkets()

	if err != nil {
		return err
	}
	if !resp.Success {
		fmt.Println("No error but no success either...")
		return nil
	}

	for _, r := range resp.Result {
		fmt.Printf("{\tname: %v\tbaseCurrency: %v}\n", r.Name, r.BaseCurrency)
	}
	return nil
}
