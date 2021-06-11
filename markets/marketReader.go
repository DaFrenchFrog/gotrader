package markets

import (
	"fmt"
	"time"

	"github.com/elRomano/gotrader/ftx"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

// MarketReader :
type MarketReader struct {
	MarketName string
	Data       model.MarketData
	client     ftx.Client
}

// NewReader :
func NewReader(mkt string) MarketReader {
	return MarketReader{
		MarketName: mkt,
		client:     ftx.Client{},
	}
}

//Load :
func (c *MarketReader) Load() error {
	err := c.loadMarketData()
	if err != nil {
		return err
	}
	return c.GetMarketHistory()
}

//GetMarketHistory :
func (c *MarketReader) GetMarketHistory() error {
	startDateToLoad := time.Now().AddDate(0, 0, -1)
	y, m, d := startDateToLoad.Date()
	cfmt.Println(cfmt.Purple, "Starting history loading : ", cfmt.Neutral, "starting date ", d, m, y)
	fmt.Println("Starting time : ", startDateToLoad)
	return c.getFramedHistory(startDateToLoad.Unix())
}

//GetLatestCandle :
func (c *MarketReader) GetLatestCandle() error {

	resp, err := c.client.GetMarketHistory(c.MarketName, 60, time.Now().Add(-time.Minute*2).Unix(), time.Now().Unix())
	if err != nil {
		return err
	}

	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error...")
		return nil
	}
	if c.Data.History[len(c.Data.History)-1].StartTime.Before(resp.Result[len(resp.Result)-1].StartTime) {
		cfmt.Println(cfmt.Cyan, "New candle ! ", cfmt.Neutral, c.Data.History[len(c.Data.History)-1].StartTime, " - ", resp.Result[len(resp.Result)-1].StartTime)
		c.Data.History = append(c.Data.History, resp.Result[len(resp.Result)-1])
	} else {
		cfmt.Println("No new candle... ", c.Data.History[len(c.Data.History)-1].StartTime, " <> ", resp.Result[len(resp.Result)-1].StartTime)
	}
	// cfmt.Println(cfmt.Cyan, "Last candle : ", cfmt.Neutral, resp.Result)
	return nil
}

//getFramedCoinHistory :
func (c *MarketReader) getFramedHistory(dateStart int64) error {
	endDate := dateStart + int64(1500*60)

	resp, err := c.client.GetMarketHistory(c.MarketName, 60, dateStart, endDate)
	if err != nil {
		return err
	}

	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error...")
		return nil
	}
	resultLength := len(resp.Result)
	cfmt.Println(cfmt.Green, "Tickers loaded : ", cfmt.Neutral, resultLength, " ", c.MarketName, " entries from \t", resp.Result[0].StartTime.Format("2 Jan 2006 15:04"), " to \t", resp.Result[resultLength-1].StartTime.Format("2 Jan 2006 15:04"))

	c.appendHistory(resp.Result)
	if endDate < time.Now().Unix() {
		c.getFramedHistory(endDate)
	}
	return nil
}

func (c *MarketReader) appendHistory(history []model.Candle) {
	c.Data.History = append(c.Data.History, history...)
}

//loadMarketData :
func (c *MarketReader) loadMarketData() error {
	resp, err := c.client.ListMarket(c.MarketName)
	if err != nil {
		return err
	}
	if !resp.Success {
		fmt.Println("No error but no success either...")
		return nil
	}

	cfmt.Println(cfmt.Green, "Market loaded: ", cfmt.Neutral, resp.Result.Name, "\tlast: ", resp.Result.Last, "\task:", resp.Result.Ask, "\tbid:", resp.Result.Bid)
	c.Data = resp.Result
	return nil
}
