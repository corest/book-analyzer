package orderbook

type OrderType int

const (
	Undefined OrderType = iota
	BidOrder
	AskOrder
)

// Order represents order part with price
type Order struct {
	Timestamp string // not used in program as time so no need in actual time type
	ID        string
	Type      OrderType
	Price     float64
}

// OrderState represents order part with shares
type OrderState struct {
	Type     OrderType
	IsActive bool
	Shares   int
}

// OrderBook represents struct where all business logic is bounded to
type OrderBook struct {
	bidShareSum int
	askShareSum int
	bids        []*Order
	asks        []*Order
	orderIDs    map[string]OrderState
	targetSize  int
}

// OrderResult represents data from processing single input
type OrderResult struct {
	Total     float64
	OrderCode string
	Timestamp string
}
