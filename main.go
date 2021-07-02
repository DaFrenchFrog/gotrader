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

	_ "net/http/pprof"
)

func main() {

	cfmt.Println(cfmt.Cyan, "\nStarting program ||||||||||||||||||||||||||||||||||||||||||||||")
	marketList := []string{"BTC/USDT" /*, "ETH/USDT", "DEFI-PERP" ,"SHIT-PERP", "ALT-PERP", "BULL/USDT", "BEAR/USDT", "ETHBULL/USDT", "ETHBEAR/USDT"*/}
	backtestCmd := flag.NewFlagSet("backtest", flag.ExitOnError)
	backtestSpeedCmd := backtestCmd.String("term", "long", "'day', 'month','year','all'")
	liveCmd := flag.NewFlagSet("live", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Fatal(model.Color("red"), "Missing command: list, backtest, updatedb, showdb, or live", model.Color(""))
	}

	fakeWallet := account.New()
	fakeWallet.RegisterMarkets(marketList, 1000)

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
	case "updatedb":
		cfmt.Println(cfmt.Yellow, "Launching database update...")
		err = store.UpdateDb(marketList)
		if err != nil {
			log.Fatal("Error updating database : ", err)
		}
	case "showdb":
		cfmt.Println(cfmt.Yellow, "Listing database content...")
		boltdb.ShowBucketContent(db)
	case "list":
		cfmt.Println(cfmt.Yellow, "Listing markets...")
		marketList := markets.NewLister()
		err = marketList.ListMarkets()
	case "backtest":
		cfmt.Println(cfmt.Yellow, "Launching backtest on ", marketList, "...")
		_ = backtestCmd.Parse(os.Args[2:])
		runner := strategy.New(marketList, strategy.NewNormalStrategy(), fakeWallet, store)
		runner.LaunchBacktest(*backtestSpeedCmd)
	case "live":
		cfmt.Println(cfmt.Yellow, "Let's lose some money baby !")
		_ = liveCmd.Parse(os.Args[2:])
		runner := strategy.New(marketList, strategy.NewNormalStrategy(), fakeWallet, store)
		runner.Live(marketList)
	default:
		fmt.Println("command unknown")
	}
	cfmt.Println(cfmt.Cyan, "|||||||||||||||||||||||||||||||||||||||||||||| Program terminated.\n")

	if err != nil {
		log.Fatal(err)
	}
}
