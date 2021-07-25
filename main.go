package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/corest/bookanalyzer/pkg/orderbook"
)

func main() {
	targetSize := flag.Int("target-size", 0, "Target size for trading")
	flag.Parse()

	if *targetSize <= 0 {
		panic("-target-size must be > 0")
	}

	scanner := bufio.NewScanner(os.Stdin)
	orderBook := orderbook.New(*targetSize)
	for scanner.Scan() {
		inputString := scanner.Text()

		result, err := orderBook.Parse(inputString)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if result != "" {
			fmt.Println(result)
		}
	}
}
