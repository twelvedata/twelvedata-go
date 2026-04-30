package main

import (
	"context"
	"fmt"
	"log"

	"github.com/twelvedata/twelvedata-go/twelvedata"
)

func main() {
	cfg, err := twelvedata.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	client := twelvedata.NewAPIClient(cfg)

	resp, _, err := client.CurrenciesAPI.GetExchangeRate(context.Background()).
		Symbol("USD/EUR").
		Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
