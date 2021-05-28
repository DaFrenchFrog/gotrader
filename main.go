package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elRomano/gotrader/coinReader"
)

func main() {
	fmt.Println("––––––––––––––––––– LET'S RIDE THE CANDLES ! –––––––––––––––––––")
	if len(os.Args) > 1 {
		reader := coinReader.New()
		if os.Args[1] == "list" {
			reader.ListMarkets("ETH/USDT")
		} else if os.Args[1] == "backtest" {
			reader.ListCoin(os.Args[2])
		}
	} else {
		log.Fatal("Missing a pair argument ! ")

	}

}
