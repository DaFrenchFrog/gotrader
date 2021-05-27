package main

import (
	"log"
	"os"

	"github.com/elRomano/gotrader/coinReader"
)

func main() {
	if len(os.Args) > 1 {
		reader := coinReader.New()
		if os.Args[1] == "list" {
			reader.ListMarkets()
		} else {
			reader.ListCoin("ETH/USDT")
		}
	} else {
		log.Fatal("Missing a pair argument ! ")

	}

}
