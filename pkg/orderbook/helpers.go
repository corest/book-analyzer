package orderbook

import (
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func floatToFixedSign(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func removeIndex(index int, data []*Order) []*Order {
	ret := make([]*Order, 0)
	ret = append(ret, data[:index]...)
	return append(ret, data[index+1:]...)
}
