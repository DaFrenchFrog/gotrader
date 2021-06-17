package markets

import (
	"fmt"
	"github.com/elRomano/gotrader/store"
	"time"

	"github.com/elRomano/gotrader/ftx"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

// MarketReader :
type MarketReader struct {
	MarketName       string
	Data             model.MarketData
	client           ftx.Client
	newCandleChannel chan model.Candle
	store            store.HistoryStore
}

// NewReader :
func NewReader(mkt string, store store.HistoryStore) MarketReader {
	return MarketReader{
		MarketName:       mkt,
		client:           ftx.Client{},
		newCandleChannel: make(chan model.Candle),
		store:            store,
	}
}

//Load :
func (m *MarketReader) Load() error {
	err := m.loadMarketData()
	if err != nil {
		return err
	}
	return m.GetMarketHistory()
}

//GetMarketHistory :
func (m *MarketReader) GetMarketHistory() error {
	startDateToLoad := time.Now().AddDate(0, -1, 0)
	cfmt.Println(cfmt.Purple, "Starting history loading : ", cfmt.Neutral, "starting date ", startDateToLoad.Format("2 Jan 2006"))
	return m.getFramedHistory(startDateToLoad.Unix())
}

//GetLatestCandle :
func (m *MarketReader) GetLatestCandle() error {
	resp, err := m.client.GetMarketHistory(m.MarketName, 60, time.Now().Add(-time.Minute*2).Unix(), time.Now().Unix())
	if err != nil {
		return err
	}

	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error...")
		return nil
	}
	if m.Data.History[len(m.Data.History)-1].StartTime.Before(resp.Result[len(resp.Result)-1].StartTime) {
		cfmt.Println(cfmt.Cyan, "New candle ! ", cfmt.Neutral, m.Data.History[len(m.Data.History)-1].StartTime, " - ", resp.Result[len(resp.Result)-1].StartTime)
		m.Data.History = append(m.Data.History, resp.Result[len(resp.Result)-1])
		m.newCandleChannel <- m.Data.History[len(m.Data.History)-1]
	} else {
		cfmt.Println("No new candle... ", m.Data.History[len(m.Data.History)-1].StartTime, " <> ", resp.Result[len(resp.Result)-1].StartTime)
	}
	// cfmt.Println(cfmt.Cyan, "Last candle : ", cfmt.Neutral, resp.Result)
	return nil
}

//NewCandleChannel :
func (m *MarketReader) GetNewCandleChannel() <-chan model.Candle {
	return m.newCandleChannel
}

//getFramedCoinHistory :
func (m *MarketReader) getFramedHistory(dateStart int64) error {
	endDate := dateStart + int64(1500*60)

	resp, err := m.client.GetMarketHistory(m.MarketName, 60, dateStart, endDate)
	if err != nil {
		return err
	}

	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error...")
		return nil
	}
	resultLength := len(resp.Result)
	cfmt.Println(cfmt.Green, "Tickers loaded : ", cfmt.Neutral, resultLength, " ", m.MarketName, " entries from \t", resp.Result[0].StartTime.Format("2 Jan 2006 15:04"), " to \t", resp.Result[resultLength-1].StartTime.Format("2 Jan 2006 15:04"))

	err = m.store.SaveBatch(resp.Result)
	if err != nil {
		cfmt.Printf(cfmt.Red, "oops can't save because, %v", err)
	}

	m.appendHistory(resp.Result)
	if endDate < time.Now().Unix() {
		m.getFramedHistory(endDate)
	}
	return nil
}

func (m *MarketReader) appendHistory(history []model.Candle) {
	m.Data.History = append(m.Data.History, history...)
}

//loadMarketData :
func (m *MarketReader) loadMarketData() error {
	resp, err := m.client.ListMarket(m.MarketName)
	if err != nil {
		return err
	}
	if !resp.Success {
		fmt.Println("No error but no success either...")
		return nil
	}

	cfmt.Println(cfmt.Green, "Market loaded: ", cfmt.Neutral, resp.Result.Name, "\tlast: ", resp.Result.Last, "\task:", resp.Result.Ask, "\tbid:", resp.Result.Bid)
	m.Data = resp.Result
	return nil
}
