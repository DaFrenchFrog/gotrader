package boltdb

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/ftx"
	"github.com/elRomano/gotrader/model"
)

type historyDataStore struct {
	db     *bolt.DB
	bucket string
	client ftx.Client
}

const mainBucket = "marketsData"

const dateOrigin = "13-09-2018 15:15:00"
const dateLayout = "02-01-2006 15:04:05"

//SetHistoryDataStore :
func SetHistoryDataStore(db *bolt.DB) (*historyDataStore, error) {
	store := &historyDataStore{
		db:     db,
		bucket: mainBucket,
		client: ftx.Client{},
	}
	return store, nil
}

func (s *historyDataStore) UpdateDb(marketList []string) error {
	s.createRootIfNotExists()
	for _, mktName := range marketList {
		fmt.Println("looking for oldest ", mktName)
		err := s.updateMarketHistory(mktName)
		if err != nil {
			return err
		}
	}
	ShowBucketContent(s.db)
	return nil
}

func (s historyDataStore) updateMarketHistory(mkt string) error {
	startDateToLoad := time.Time{}
	if !s.MarketExist(mkt) {
		cfmt.Println(cfmt.Blue, "New market detected : ", mkt)
		s.initNewMarket(mkt)
		startDateToLoad = s.getOldestCandleDate(mkt)
	} else {
		startDateToLoad = s.getLatestDateKey(mkt)
	}

	cfmt.Println(cfmt.Blue, "Starting update : ", cfmt.Neutral, mkt, " from ", startDateToLoad.Format(model.DateLayoutLog), "...")

	err := s.saveCandlesRecursive(mkt, startDateToLoad.Unix())

	return err
}

func (s historyDataStore) initNewMarket(mkt string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.Bucket([]byte(mainBucket)).CreateBucketIfNotExists([]byte(mkt))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s historyDataStore) putCandlesToDb(mkt string, candles []model.Candle) error {

	tx, err := s.db.Begin(true)
	b := tx.Bucket([]byte(mainBucket)).Bucket([]byte(mkt))
	i := 0

	for _, candle := range candles {
		i++

		js, err := json.Marshal(candle)
		if err != nil {
			return err
		}

		err = b.Put(itob(candle.StartTime.Unix()), js)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	cfmt.Print(cfmt.Green, " ...and saved : ", cfmt.Neutral, candles[0].StartTime.Format(model.DateLayoutLog), "\t to \t", candles[len(candles)-1].StartTime.Format(model.DateLayoutLog), " on ", mkt, " market\n")
	return err
}

//getFramedCoinHistory :
func (s historyDataStore) saveCandlesRecursive(mkt string, dateStart int64) error {
	endDate := dateStart + int64(1500*60)
	resp, err := s.client.GetMarketHistory(mkt, 60, dateStart, endDate)

	if err != nil {
		return err
	}
	if !resp.Success {
		cfmt.Println(cfmt.Red, "There was an error retrieving ", mkt, " market history... ")
		log.Fatal("!!!!!!! Error !!!!!! history length = ", len(resp.Result), "\tsuccess = ", resp.Success)
		return nil
	}
	if err != nil {
		cfmt.Printf(cfmt.Red, "oops can't save because, %v", err)
	}
	if len(resp.Result) == 0 {
		cfmt.Println(cfmt.Blue, "No data retrieved : ", cfmt.Neutral, " start : ", time.Unix(dateStart, 0).Format(model.DateLayoutLog), "(", dateStart, ")", "   end : ", time.Unix(endDate, 0).Format(model.DateLayoutLog), "(", endDate, ")")
		return nil
	}
	s.logLoadedFrame(mkt, resp.Result)
	// candles = append(candles, resp.Result...)
	err = s.putCandlesToDb(mkt, resp.Result)
	if err != nil {
		return err
	}
	if endDate < time.Now().Unix() {
		s.saveCandlesRecursive(mkt, endDate)
	}
	return nil
}

func (s historyDataStore) logLoadedFrame(mkt string, resp []model.Candle) {
	resultLength := len(resp)
	cfmt.Print(cfmt.Green, resultLength, " Tickers loaded... ")
}

func (s historyDataStore) getOldestCandleDate(mkt string) time.Time {
	searchInterval := 3600
	d, _ := time.Parse(dateLayout, dateOrigin)
	dateStart := d.Unix()
	oldestDate := time.Time{}
	for {
		endDate := dateStart + int64(1500*searchInterval)
		resp, err := s.client.GetMarketHistory(mkt, int64(searchInterval), dateStart, endDate)
		// fmt.Println("Looking for a candle : ", time.Unix(dateStart, 0))
		dateStart = endDate
		if err != nil {
			log.Fatal("Error retrieving candles : ", err)
		}
		if len(resp.Result) > 0 {
			// fmt.Println("break loop", len(resp.Result), resp.Result[0].StartTime.Format(model.DateLayoutLog))
			oldestDate = resp.Result[0].StartTime
			break
		}
	}
	// fmt.Println("resp length = ", len(resp.Result))
	return oldestDate
}

func (s historyDataStore) getLatestDateKey(mkt string) time.Time {
	tx, err := s.db.Begin(false)
	if err != nil {
		log.Fatal("Error getting latest Date Key in DB")
	}
	b := tx.Bucket([]byte(mainBucket)).Bucket([]byte(mkt))

	k, v := b.Cursor().Last()
	t := time.Time{}
	// fmt.Println(" k  = ", time.Unix(btoi(k), 0))
	if string(v) != "" {
		t = time.Unix(btoi(k), 0)
	} else {
		t, _ = time.Parse(dateLayout, dateOrigin)
	}
	err = tx.Rollback()
	if err != nil {
		log.Fatal("Rollback Error getting latest Date Key in DB")
	}
	return t
}

// func (s historyDataStore) Save(candle model.Candle) error {
// 	err := s.db.Update(func(tx *bolt.Tx) error {
// 		return s.saveCandle(tx, candle)
// 	})

// 	return err
// }
func (s historyDataStore) GetMarketHistory(bucket string, term string) ([]model.Candle, error) {
	result := []model.Candle{}
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(mainBucket)).Bucket([]byte(bucket))
		c := b.Cursor()
		key, _ := c.Seek(s.getStartingKeyToLoad(b, term))
		cfmt.Println(cfmt.Blue, "since ", time.Unix(btoi(key), 0).Format(model.DateLayoutLog))
		for k, v := c.Seek(key); k != nil; k, v = c.Next() {
			// fmt.Printf("key=%s, value=%s\n", k, v)
			// }
			// err := b.ForEach(func(k, v []byte) error {
			md := model.Candle{}
			err := json.Unmarshal(v, &md)

			if err != nil {
				fmt.Println("EEERRERRRk=", k)
				return err
			}
			result = append(result, md)

		}

		// if err != nil {
		// 	return err
		// }

		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s historyDataStore) getStartingKeyToLoad(b *bolt.Bucket, term string) []byte {
	candleAmount := int64(0)
	switch term {
	case "day":
		//1 day in seconds
		candleAmount = 60 * 60 * 24
	case "month":
		//1 month in seconds
		candleAmount = 60 * 60 * 24 * 30
	case "year":
		//1 year in seconds
		candleAmount = 60 * 60 * 24 * 365
	default:
		//all data
		candleAmount = 0
	}
	c := b.Cursor()
	k, _ := c.Last()
	if candleAmount == 0 {
		k, _ = c.First()
	}
	n := btoi(k)
	return []byte(itob(n - candleAmount))
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

//MarketExist :
func (s historyDataStore) MarketExist(mkt string) bool {
	tx, err := s.db.Begin(false)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	b := tx.Bucket([]byte(mainBucket)).Bucket([]byte(mkt))
	defer tx.Rollback()
	return b != nil
}

// createRootIfNotExists :
func (s historyDataStore) createRootIfNotExists() error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(mainBucket))
		if err != nil {
			return fmt.Errorf("Error on creating main bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

//ShowBucketContent :
func ShowBucketContent(db *bolt.DB) {

	cfmt.Printf(cfmt.Blue, "Showing database content...")
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(mainBucket))

		b.ForEach(func(k, v []byte) error {

			subB := b.Bucket(k)
			f, _ := subB.Cursor().First()
			l, _ := subB.Cursor().Last()
			if len(l) == 0 || len(f) == 0 {
				cfmt.Println(cfmt.Green, "Market ", string(k), cfmt.Neutral, " => Empty !")
				return nil
			}
			firstDate := time.Unix(btoi(f), 0)
			lastDate := time.Unix(btoi(l), 0)
			nbEntries := 0
			if subB != nil {
				subB.ForEach(func(l, w []byte) error {
					// fmt.Println(l, "\t", btoi(l), "\t", time.Unix(btoi(l), 0).Format(model.DateLayoutLog))
					nbEntries++
					return nil
				})
			}
			cfmt.Println(cfmt.Green, "Martket ", string(k), cfmt.Neutral, "\t => \t", nbEntries, " candles from \t", firstDate.Format(model.DateLayoutLog), " to \t", lastDate.Format(model.DateLayoutLog))
			return nil
		})
		return nil
	})
	if err != nil {
		cfmt.Println(cfmt.Red, err)
	}
}

// itob returns an 8-byte big endian representation of v.
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}
