package orderbook

type OrderSide int
type OperationType int
type TradeAction int

const (
	Undefined OrderSide = iota
	BidOrder
	AskOrder
)

const (
	UnknownOp OperationType = iota
	AddOrder
	ReduceOrder
)

const (
	Buy TradeAction = iota
	Sell
)

// Order represents order part with price
type Order struct {
	Timestamp string // not used in program as time so no need in actual time type
	ID        string
	Side      OrderSide
	Price     float64
	Size      int
	Operation OperationType
}

// OrderBook represents struct where all business logic is bounded to
type OrderBook struct {
	orders       map[string]Order
	totalShares  map[OrderSide]int
	orderIDs     map[string]OrderSide
	canTrade     map[TradeAction]bool
	totalHistory map[TradeAction]float64
	targetSize   int
}
