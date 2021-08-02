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

func floatEqual(a, b float64) bool {
	ai, bi := int64(math.Float64bits(a)), int64(math.Float64bits(b))
	return a == b || -1 <= ai-bi && ai-bi <= 1
}

func formatResult(total float64, timestamp, tradeAction string) string {
	if total > 0.0 {
		return fmt.Sprintf("%s %s %.2f", timestamp, tradeAction, total)
	} else {
		return fmt.Sprintf("%s %s NA", timestamp, tradeAction)
	}
}
