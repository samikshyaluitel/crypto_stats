package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/crypto_stats/handler"
	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	symbols = kingpin.Arg("symbols", "Comma separated symbols").Default("BTCUSD,ETHBTC").String()
)

func main() {
	kingpin.Parse()
	supportedSymbols := strings.Split(*symbols, ",")
	fmt.Printf("Supported Symbols %v",supportedSymbols)
	ca, err := handler.NewCurrencyAPI(http.DefaultClient, supportedSymbols)
	if err != nil {
		panic(err)
	}

	router := echo.New()
	router.GET("/currency/:symbol", ca.GetCurrency)

	log.Fatal(http.ListenAndServe(":8080", router))
}
