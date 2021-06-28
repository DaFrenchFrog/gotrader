package store

import "github.com/elRomano/gotrader/model"

type HistoryStore interface {
	GetMarketHistory(bucket string) ([]model.Candle, error)
	MarketExist(bucket string) bool
	UpdateDb(marketList []string) error
}
