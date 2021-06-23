package store

import "github.com/elRomano/gotrader/model"

type HistoryStore interface {
	Save(model.Candle) error
	SaveBatch([]model.Candle) error
	GetAll() ([]model.Candle, error)
	GetMarket(bucket string) ([]model.Candle, error)
	MarketExist(bucket string) bool
}
