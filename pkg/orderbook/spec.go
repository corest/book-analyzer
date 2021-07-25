package orderbook

type OrderType int
type TradingAction int

const (
	Undefined OrderType = iota
	BidOrder
	AskOrder
)

const (
	Buy TradingAction = iota
	Sell
)

type Order struct {
	Timestamp string // not used in program as time so no need in actual time type
	ID        string
	Type      OrderType
	Price     float64
}

type OrderState struct {
	Type     OrderType
	IsActive bool
	Shares   int
}

type OrderBook struct {
	bidShareSum int
	askShareSum int
	bids        []*Order
	asks        []*Order
	orderIDs    map[string]OrderState
	targetSize  int
}

type OrderExecution struct {
	Timestamp string
	Action    TradingAction
	Total     string
}

type OrderResult struct {
	Total     float64
	OrderCode string
	Timestamp string
}
