package coinreader

import (
	"fmt"
	"time"

	"github.com/elRomano/gotrader/apiCaller"
	"github.com/elRomano/gotrader/cfmt"
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
func (c *CoinReader) GetCoinData(coin string) error {
	c.listCoin(coin)
	c.GetCoinHistory(coin)
	return nil
}

//GetCoinHistory :
func (c *CoinReader) GetCoinHistory(coin string) error {
	startDateToLoad := time.Now().AddDate(0, -1, 0)
	y, m, d := startDateToLoad.Date()
	cfmt.Println(cfmt.Purple, "Starting history loading : ", cfmt.Neutral, "starting date ", d, m, y)
	fmt.Println("Starting time : ", startDateToLoad)
	c.getFramedCoinHistory(coin, startDateToLoad.Unix())
	return nil
}

//getFramedCoinHistory :
func (c *CoinReader) getFramedCoinHistory(coin string, dateStart int64) error {
	timeframe := 60
	endDate := dateStart + int64(1500*timeframe)

	resp := &model.CoinHistoryResponse{}
	succeed, err := apiCaller.Call(baseURL+"/markets/"+coin+"/candles?resolution=60&start_time="+fmt.Sprint(dateStart)+"&end_time="+fmt.Sprint(endDate), resp)
	if err != nil {
		return err
	}
	if !succeed {
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

//ListCoin :
func (c *CoinReader) listCoin(coin string) error {
	resp := &model.CoinDataResponse{}
	succeed, err := apiCaller.Call(baseURL+"/markets/"+coin, resp)

	if err != nil {
		return err
	}
	if !succeed {
		fmt.Println("No error but no success either...")
		return nil
	}

	fmt.Println(model.Color("green"), "Market loaded: ", model.Color(""), " name: ", resp.Result.Name, "\tbaseCurrency: ", resp.Result.BaseCurrency)
	c.Market = resp.Result
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
