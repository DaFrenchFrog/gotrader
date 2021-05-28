package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/elRomano/gotrader/coinReader"
)

func main() {
	// les flags te permet de passer des parametre  du style -cur=ZECBULL/USD
	// ici je crée un flagset pour la commande read
	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	// pour la commande read je met une option -cur pour preciser la currency
	readCur := readCmd.String("cur", "ETH/USDT", "The currency to read, default:ETH/USD ")

	//J'ai inverser le if: fail fast et evite de faire un else
	if len(os.Args) < 2 {
		log.Fatal("Missing command: list or read ")
	}

	reader := coinReader.New()
	var err error

	switch os.Args[1] {
	case "list":
		err = reader.ListMarkets()
	case "read":
		_ = readCmd.Parse(os.Args[2:])
		err = reader.ListCoin(*readCur)
	default:
		fmt.Println("command unknown")
	}

	if err != nil {
		log.Fatal(err)
	}
}
