package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/corest/bookanalyzer/pkg/orderbook"
)

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	orderBook := orderbook.New()
	for scanner.Scan() {
		inputString := scanner.Text()
		output, err := orderBook.Parse(inputString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if output != "" {
			fmt.Println(output)
		}
	}

	orderBook.ShowBids()
	orderBook.ShowAsks()
	orderBook.ShowStates()
}
