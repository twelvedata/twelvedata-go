package main

import (
	"context"
	"errors"
	"log"

	"github.com/twelvedata/twelvedata-go/twelvedata/ws"
)

func main() {
	client, err := ws.NewClient(ws.Options{}) // APIKey defaults to $TWELVEDATA_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	go func() {
		for p := range client.Prices() {
			log.Printf("%s @ %v", p.Symbol, p.Price)
		}
	}()
	go func() {
		for s := range client.SubscribeStatuses() {
			log.Printf("subscribe-status: %s, success=%d, fails=%d", s.Status, len(s.Success), len(s.Fails))
		}
	}()
	go func() {
		for r := range client.Reconnecting() {
			log.Printf("reconnecting: attempt=%d, in=%s", r.Attempt, r.Delay)
		}
	}()
	go func() {
		for err := range client.Errors() {
			var authErr *ws.AuthError
			if errors.As(err, &authErr) {
				log.Fatalf("auth: %v", authErr)
			}
			log.Printf("ws error: %v", err)
		}
	}()

	if err := client.Connect(context.Background()); err != nil {
		var authErr *ws.AuthError
		if errors.As(err, &authErr) {
			log.Fatalf("auth: %v", authErr)
		}
		log.Fatal(err)
	}

	if err := client.Subscribe("AAPL,EUR/USD"); err != nil {
		log.Fatal(err)
	}

	select {} // run forever
}
