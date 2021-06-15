package boltdb

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/elRomano/gotrader/model"
)

type historyDataStore struct {
	db     *bolt.DB
	bucket string
}

func NewHistoryDataStore(db *bolt.DB, bucket string) (*historyDataStore, error) {
	store := &historyDataStore{
		db:     db,
		bucket: bucket,
	}

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s historyDataStore) Save(candle model.Candle) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		return s.saveCandle(tx, candle)
	})

	return err
}

func (s historyDataStore) SaveBatch(candles []model.Candle) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		for _, candle := range candles {
			return s.saveCandle(tx, candle)
		}
		return nil
	})

	return err
}

func (s historyDataStore) saveCandle(tx *bolt.Tx, candle model.Candle) error {
	b := tx.Bucket([]byte(s.bucket))

	key, err := b.NextSequence()
	if err != nil {
		return err
	}

	js, err := json.Marshal(candle)
	if err != nil {
		return err
	}

	err = b.Put(itob(key), js)
	return err
}

func (s historyDataStore) GetAll() ([]model.Candle, error) {
	result := make([]model.Candle, 0)
	err := s.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(s.bucket))

		err := b.ForEach(func(k, v []byte) error {
			md := model.Candle{}
			err := json.Unmarshal(v, &md)
			if err != nil {
				return err
			}
			result = append(result, md)

			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
