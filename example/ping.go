package main

import (
	"fmt"
	"log"

	gecko "github.com/neox5/go-coingecko/v3"
)

func main() {
	resp, err := gecko.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.GeckoSays)
}
