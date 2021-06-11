package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/cfmt"
	"github.com/elRomano/gotrader/markets"
	"github.com/elRomano/gotrader/model"
	"github.com/elRomano/gotrader/strategy"
)

func main() {
	readCmd := flag.NewFlagSet("backtest", flag.ExitOnError)
	readMkt := readCmd.String("mkt", "ETH/USDT", "The market to read, default:ETH/USD ")

	liveCmd := flag.NewFlagSet("live", flag.ExitOnError)
	liveMkt := liveCmd.String("mkt", "ETH/USDT", "The market to trade with, default:ETH/USD ")

	if len(os.Args) < 2 {
		log.Fatal(model.Color("red"), "Missing command: list or backtest", model.Color(""))
	}

	// reader := coinreader.New()
	//runner := strategy.New(strategy.CrazyStrategy{}) // example de different strategy avec une interface

	var err error

	switch os.Args[1] {
	case "list":
		marketList := markets.NewLister()
		err = marketList.ListMarkets()
	case "backtest":
		_ = readCmd.Parse(os.Args[2:])
		runner := strategy.New(strategy.NewNormalStrategy(), account.Wallet{})
		runner.Backtest(*readMkt)
		cfmt.Println(cfmt.Cyan, "|||||||||||||||||||||||||||||||||||||||||||||| Program terminated.")
	case "live":
		_ = liveCmd.Parse(os.Args[2:])
		runner := strategy.New(strategy.NewNormalStrategy(), account.Wallet{})
		runner.Live(*liveMkt)

	default:
		fmt.Println("command unknown")
	}

	if err != nil {
		log.Fatal(err)
	}
}
