package orderbook

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type OrderType int

const (
	Undefined OrderType = iota
	BidOrder
	AskOrder
)

type Order struct {
	Timestamp string // not used in program widely so no need in actual time
	ID        string
	Shares    int
	Type      OrderType
	Price     float64
}

type OrderBook struct {
	bids     []*Order
	asks     []*Order
	orderIDs map[string]bool
}

func New() *OrderBook {
	return &OrderBook{
		orderIDs: map[string]bool{},
	}
}

func (ob *OrderBook) Parse(inputString string) (string, error) {
	in := strings.Split(inputString, " ")

	// add/remove order always has the same amount of shares for particular ID
	// therefore it is either added or removed
	// there is no need to have update order
	switch len(in) {
	case 6:
		// format: "timestamp operation id order_type price shares"
		price, _ := strconv.ParseFloat(in[4], 32) // ignoring error for now
		shares, _ := strconv.Atoi(in[5])
		orderType := getOrderType(in[3])
		newOrder := &Order{
			Timestamp: in[0],
			ID:        in[2],
			Type:      orderType,
			Price:     floatToFixedSign(price, 2),
			Shares:    shares,
		}
		return ob.addOrder(newOrder)
	case 4:
		// format: "timestamp operation id shares"
		shares, _ := strconv.Atoi(in[3])
		return ob.markOrderDeleted(in[2], shares)
	default:
		return "", nil // error failed to parse and continue
	}
}

func (ob *OrderBook) addOrder(order *Order) (string, error) {
	if order.Type == Undefined {
		return "", errors.New("unknown order type")
	}

	ob.orderIDs[order.ID] = true

	if order.Type == BidOrder {
		return ob.addBidOrder(order)
	}

	return ob.addAskOrder(order)
}

func (ob *OrderBook) addBidOrder(order *Order) (string, error) {
	ob.bids = addSortedOrder(order, ob.bids)
	return "", nil
}

func (ob *OrderBook) addAskOrder(order *Order) (string, error) {
	ob.asks = addSortedOrder(order, ob.asks)
	return "", nil
}

func (ob *OrderBook) markOrderDeleted(id string, shares int) (string, error) {
	ob.orderIDs[id] = false

	return "", nil
}

// debug purpose
func (ob *OrderBook) ShowBids() {
	for i, o := range ob.bids {
		if ob.orderIDs[o.ID] {
			fmt.Printf("bid | id: %s | %d: %f\n", o.ID, i, o.Price)
		}
	}
}

func (ob *OrderBook) ShowAsks() {
	for i, o := range ob.asks {
		if ob.orderIDs[o.ID] {
			fmt.Printf("ask | id %s | %d: %f\n", o.ID, i, o.Price)
		}
	}
}

func (ob *OrderBook) ShowStates() {
	fmt.Printf("\n%v\n", ob.orderIDs)
}

func getOrderType(str string) OrderType {
	switch str {
	case "B":
		return BidOrder
	case "S":
		return AskOrder
	default:
		return Undefined
	}
}

func addSortedOrder(order *Order, data []*Order) []*Order {
	i := sort.Search(len(data), func(i int) bool {
		return data[i].Price >= order.Price
	})

	if i == len(data) {
		return append(data, order)
	}

	// make space for the inserted element by shifting values
	data = append(data[:i+1], data[i:]...)
	data[i] = order
	return data
}
