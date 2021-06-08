package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/elRomano/gotrader/cfmt"
	coinreader "github.com/elRomano/gotrader/coinReader"
	"github.com/elRomano/gotrader/model"
	"github.com/elRomano/gotrader/strategy"
)

func main() {
	readCmd := flag.NewFlagSet("backtest", flag.ExitOnError)
	readCur := readCmd.String("mkt", "ETH/USDT", "The market to read, default:ETH/USD ")

	if len(os.Args) < 2 {
		log.Fatal(model.Color("red"), "Missing command: list or backtest", model.Color(""))
	}

	reader := coinreader.New()
	strategy := strategy.New()

	var err error

	switch os.Args[1] {
	case "list":
		err = reader.ListMarkets()
	case "backtest":
		_ = readCmd.Parse(os.Args[2:])
		err = reader.GetCoinData(*readCur)
		strategy.Backtest(reader.Market)
		cfmt.Println(cfmt.Cyan, "Program terminated.")
	default:
		fmt.Println("command unknown")
	}

	if err != nil {
		log.Fatal(err)
	}
}
