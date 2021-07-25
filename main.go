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
	var previousBuyResult, previousSellResult float64
	for scanner.Scan() {
		inputString := scanner.Text()

		result, err := orderBook.Parse(inputString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if result.OrderCode == "B" && orderbook.FloatEqual(previousBuyResult, result.Total) {
			continue
		}
		if result.OrderCode == "S" && orderbook.FloatEqual(previousSellResult, result.Total) {
			continue
		}

		output := orderbook.FormatResult(result)

		if output != "" {
			if result.OrderCode == "B" {
				previousBuyResult = result.Total
			} else {
				previousSellResult = result.Total

			}
			fmt.Println(output)
		}
	}
}
