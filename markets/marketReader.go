package markets

import (
	"fmt"
	"log"

	"github.com/elRomano/gotrader/store"

	"github.com/elRomano/gotrader/ftx"

	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/model"
)

// MarketReader :
type MarketReader struct {
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
		cfmt.Println(cfmt.Blue, "Loading ", mData.Name, " market...")
		m.loadMarketSummary(mData.Name)
		cfmt.Println(cfmt.Blue, "Market summary loaded. Loading history...", cfmt.Neutral)
		if m.store.MarketExist(mData.Name) {
			m.Data[mData.Name].History, _ = m.store.GetMarketHistory(mData.Name)
		} else {
			cfmt.Println(cfmt.Yellow, "Market ", mData.Name, " does not exist in database. Run -update to retrieve data.")
		}
		cfmt.Println(cfmt.Green, mData.Name, " loaded : ", cfmt.Neutral, len(m.Data[mData.Name].History), " entries from ", m.Data[mData.Name].History[0].StartTime.Format(model.DateLayoutLog))
	}

	return nil
	// return m.GetMarketHistory()
}

//GetNewCandleChannel :
func (m *MarketReader) GetNewCandleChannel() <-chan model.Candle {
	return m.newCandleChannel
}

//loadMarketData :
func (m *MarketReader) loadMarketSummary(mkt string) error {
	resp, err := m.client.GetMarketSummary(mkt)
	if err != nil {
		return err
	}
	if !resp.Success {
		fmt.Println("No error but no success either...")
		return nil
	}

	if err != nil {
		log.Fatalf("Error loading MarketData : %v", err)
	}
	m.Data[mkt] = &resp.Result

	return nil
}
