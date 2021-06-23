package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/elRomano/gotrader/store/boltdb"

	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/markets"
	"github.com/elRomano/gotrader/model"
	"github.com/elRomano/gotrader/strategy"
)

func main() {
	marketList := []string{"BTC/USD", "ETH/USD"}
	readCmd := flag.NewFlagSet("backtest", flag.ExitOnError)
	// readMkt := readCmd.String("mkt", "ETH/USDT", "The market to read, default:ETH/USD ")

	liveCmd := flag.NewFlagSet("live", flag.ExitOnError)
	// liveMkt := liveCmd.String("mkt", "ETH/USDT", "The market to trade with, default:ETH/USD ")

	if len(os.Args) < 2 {
		log.Fatal(model.Color("red"), "Missing command: list or backtest", model.Color(""))
	}

	// reader := coinreader.New()
	//runner := strategy.New(strategy.CrazyStrategy{}) // example de different strategy avec une interface

	var err error

	db, err := bolt.Open("./gotrader.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("could not open bolt db file,  %v", err)
	}
	defer db.Close()
	store, err := boltdb.SetHistoryDataStore(db)
	if err != nil {
		log.Fatalf("could not create market data store,  %v", err)
	}

	switch os.Args[1] {
	case "updateDb":
		cfmt.Println(cfmt.Blue, "Launching database update...")
		//We check db and insert new candles
	case "showdb":
		cfmt.Println(cfmt.Blue, "Listing database content...")
		boltdb.ShowBucketContent(db)
	case "list":
		cfmt.Println(cfmt.Blue, "Listing markets...")
		marketList := markets.NewLister()
		err = marketList.ListMarkets()
	case "backtest":
		cfmt.Println(cfmt.Blue, "Launching backtest on ", marketList, "...")
		_ = readCmd.Parse(os.Args[2:])
		runner := strategy.New(marketList, strategy.NewNormalStrategy(), account.Wallet{}, store)
		runner.LaunchBacktest()
	case "live":
		cfmt.Println(cfmt.Blue, "Let's lose some money baby !")
		_ = liveCmd.Parse(os.Args[2:])
		runner := strategy.New(marketList, strategy.NewNormalStrategy(), account.Wallet{}, store)
		runner.Live(marketList)
	default:
		fmt.Println("command unknown")
	}
	cfmt.Println(cfmt.Cyan, "|||||||||||||||||||||||||||||||||||||||||||||| Program terminated.")

	if err != nil {
		log.Fatal(err)
	}
}
