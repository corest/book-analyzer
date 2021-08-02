package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/corest/bookanalyzer/pkg/orderbook"
)

func main() {
	start := time.Now()

	targetSize := flag.Int("target-size", 0, "Target size for trading")
	flag.Parse()

	if *targetSize <= 0 {
		panic("-target-size must be > 0")
	}

	orderBook := orderbook.New(*targetSize)

	scanner := bufio.NewScanner(os.Stdin)

	err := orderBook.Process(scanner)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s", elapsed)
}
