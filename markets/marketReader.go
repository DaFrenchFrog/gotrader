package markets

import (
	"fmt"
	"time"

	"github.com/elRomano/gotrader/store"

	"github.com/elRomano/gotrader/ftx"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

// MarketReader :
type MarketReader struct {
	// MarketNames      string
	Data             map[string]*model.MarketData
	client           ftx.Client
	newCandleChannel chan model.Candle
	store            store.HistoryStore
}

// NewReader :
func NewReader(marketNames []string, store store.HistoryStore) MarketReader {

	d := make(map[string]*model.MarketData, len(marketNames))
	for _, m := range marketNames {
		d[m] = &model.MarketData{}
		d[m].Name = m
	}
	return MarketReader{
		Data:             d,
		client:           ftx.Client{},
		newCandleChannel: make(chan model.Candle),
		store:            store,
	}
}

//Load :
func (m *MarketReader) Load() error {
	for _, mData := range m.Data {
		if m.store.MarketExist(mData.Name) {
			m.store.GetAll()
		} else {
			cfmt.Println(cfmt.Yellow, "Market ", mData.Name, " does not exist in database. Run -update to retrieve data.")
		}

		resp, err := m.client.GetMarketSummary(mData.Name)
		if err != nil {
			return err
		}
		if !resp.Success {
			fmt.Println("No error but no success either...")
			return nil
		}

		cfmt.Println(cfmt.Green, "Market loaded: ", cfmt.Neutral, resp.Result.Name, "\tlast: ", resp.Result.Last, "\task:", resp.Result.Ask, "\tbid:", resp.Result.Bid)

		if err != nil {
			fmt.Errorf("Error loading MarketData : ", err)
		}
		m.Data[mData.Name] = &resp.Result
	}

	return nil
	// return m.GetMarketHistory()
}

//GetMarketHistory :
func (m *MarketReader) GetMarketHistory(mkt string) error {
	startDateToLoad := time.Now().AddDate(0, -1, 0)
	cfmt.Println(cfmt.Purple, "Starting history loading : ", cfmt.Neutral, "starting date ", startDateToLoad.Format("2 Jan 2006"))
	return m.getFramedHistory(mkt, startDateToLoad.Unix())
}

//GetNewCandleChannel :
func (m *MarketReader) GetNewCandleChannel() <-chan model.Candle {
	return m.newCandleChannel
}

//getFramedCoinHistory :
func (m *MarketReader) getFramedHistory(mkt string, dateStart int64) error {
	endDate := dateStart + int64(1500*60)

	resp, err := m.client.GetMarketHistory(mkt, 60, dateStart, endDate)
	if err != nil {
		return err
	}

	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error...")
		return nil
	}
	resultLength := len(resp.Result)
	cfmt.Println(cfmt.Green, "Tickers loaded : ", cfmt.Neutral, resultLength, " ", mkt, " entries from \t", resp.Result[0].StartTime.Format("2 Jan 2006 15:04"), " to \t", resp.Result[resultLength-1].StartTime.Format("2 Jan 2006 15:04"))

	err = m.store.SaveBatch(resp.Result)
	if err != nil {
		cfmt.Printf(cfmt.Red, "oops can't save because, %v", err)
	}

	// m.appendHistory(resp.Result)
	m.Data[mkt].History = append(m.Data[mkt].History, resp.Result...)
	if endDate < time.Now().Unix() {
		m.getFramedHistory(mkt, endDate)
	}
	return nil
}

// func (m *MarketReader) appendHistory(history []model.Candle) {
// 	m.Data.History = append(m.Data, history...)
// }

//loadMarketData :
func (m *MarketReader) loadMarketData(mkt string) error {
	// append(m.Data, resp.Result)
	// m.Data = resp.Result

	return nil
}
