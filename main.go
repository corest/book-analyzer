package main

import (
	"flag"

	"github.com/corest/bookanalyzer/pkg/orderbook"
)

func main() {
	targetSize := flag.Int("target-size", 0, "Target size for trading")
	flag.Parse()

	if *targetSize <= 0 {
		panic("-target-size must be > 0")
	}

	orderBook := orderbook.New(*targetSize)
	err := orderBook.Process()
	if err != nil {
		panic(err)
	}
}
