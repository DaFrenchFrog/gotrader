package markets

import (
	"fmt"
	"log"

	"github.com/elRomano/gotrader/ftx"
	"github.com/elRomano/gotrader/store"

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
		d[m].History := make(map[string][]model.Candle, 3)
		fmt.Print(f)
		//candles sur 1 mn
		d[m].History["1m"] = []model.Candle{}
		//candles sur 1 h calculées à partir du 1 mn
		d[m].History["1h"] = []model.Candle{}
		//candles sur 1 journée calculée à partir du 1h
		d[m].History["1d"] = []model.Candle{}
	}
	return MarketReader{
		Data:             d,
		client:           ftx.Client{},
		newCandleChannel: make(chan model.Candle),
		store:            store,
	}
}

//Load :
func (m *MarketReader) Load(term string) error {
	for _, mData := range m.Data {
		cfmt.Println(cfmt.Blue, "[1/3] Loading ", mData.Name, " market...")
		m.loadMarketSummary(mData.Name)
		cfmt.Print(cfmt.Blue, "[2/3] Loading history ")
		if m.store.MarketExist(mData.Name) {
			m.Data[mData.Name].History["1m"], _ = m.store.GetMarketHistory(mData.Name, term)
		} else {
			cfmt.Println(cfmt.Yellow, "Market ", mData.Name, " does not exist in database. Run -update to retrieve data.")
		}
		cfmt.Println(cfmt.Blue, "[3/3] Adding indicators...", cfmt.Neutral)
		m.computeCandles()
		m.addIndicators(m.Data[mData.Name].History["1m"])
		// gfx.DrawChart(m.Data[mData.Name].History)
		cfmt.Println(cfmt.Green, mData.Name, " loaded : ", cfmt.Neutral, len(m.Data[mData.Name].History), " entries from ", m.Data[mData.Name].History["1m"][0].StartTime.Format(model.DateLayoutLog))
	}
	return nil
	// return m.GetMarketHistory()
}

func (m *MarketReader) computeCandles() {
	// calcul des 1h et 1d
}

func (m *MarketReader) addIndicators(candles []model.Candle) {
	// i := 0
	for i := range candles {
		candles[i].SMA200 = getSMA(200, candles, i)
		candles[i].ATR14 = getATR(14, candles, i)
		candles[i].ATR7 = getATR(7, candles, i)
		candles[i].STrend = getSupertrend(3, candles, i)
	}
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
