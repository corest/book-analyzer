package orderbook

import (
	"fmt"
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

func FloatEqual(a, b float64) bool {
	ai, bi := int64(math.Float64bits(a)), int64(math.Float64bits(b))
	return a == b || -1 <= ai-bi && ai-bi <= 1
}

func FormatResult(o *OrderResult) string {
	var result string

	switch t := o.Total; {
	case t > 0:
		result = fmt.Sprintf("%s %s %.2f", o.Timestamp, o.OrderCode, o.Total)
	case t < 0:
		result = fmt.Sprintf("%s %s NA", o.Timestamp, o.OrderCode)
	}

	return result
}
