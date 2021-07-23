package orderbook

import (
	"errors"
	"fmt"
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
	Type      OrderType
	Shares    int
	Sum       float64
}

type OrderBook struct {
	bids     []*Order
	asks     []*Order
	orderIDs map[string]OrderType
}

func New() *OrderBook {
	return &OrderBook{
		orderIDs: map[string]OrderType{},
	}
}

func (ob *OrderBook) Parse(inputString string) (string, error) {
	in := strings.Split(inputString, " ")

	switch len(in) {
	case 6:
		// format: "timestamp operation id order_type sum shares"
		sum, _ := strconv.ParseFloat(in[4], 32) // ignoring error for now
		shares, _ := strconv.Atoi(in[5])
		orderType := getOrderType(in[3])
		newOrder := &Order{
			Timestamp: in[0],
			ID:        in[2],
			Type:      orderType,
			Sum:       sum,
			Shares:    shares,
		}
		return ob.addOrder(newOrder)
	case 4:
		// format: "timestamp operation id shares"
		shares, _ := strconv.Atoi(in[3])
		return ob.removeOrder(in[2], shares)
	default:
		return "", nil // error failed to parse and continue
	}
}

func (ob *OrderBook) addOrder(order *Order) (string, error) {
	if order.Type == Undefined {
		return "", errors.New("unknown order type")
	}

	ob.orderIDs[order.ID] = order.Type

	if order.Type == BidOrder {
		return ob.addBidOrder(order)
	}

	return ob.addAskOrder(order)
}

func (ob *OrderBook) addBidOrder(order *Order) (string, error) {
	fmt.Printf("Bid order with id: %#q\n", order.ID)
	ob.bids = append(ob.bids, order)
	return "", nil
}

func (ob *OrderBook) addAskOrder(order *Order) (string, error) {
	fmt.Printf("Ask order with id: %#q\n", order.ID)
	ob.asks = append(ob.asks, order)
	return "", nil
}

func (ob *OrderBook) removeOrder(id string, shares int) (string, error) {
	orderType := ob.orderIDs[id]
	delete(ob.orderIDs, id)
	if orderType == BidOrder {
		fmt.Printf("Remove bid order with id: %#q\n", id)
		return ob.removeBidOrder(id, shares)
	} else {
		fmt.Printf("Remove ask order with id: %#q\n", id)
		return ob.removeAskOrder(id, shares)
	}
}

func (ob *OrderBook) removeBidOrder(id string, shares int) (string, error) {
	// remove from ob.bids
	return "", nil
}

func (ob *OrderBook) removeAskOrder(id string, shares int) (string, error) {
	// remove from ob.bids
	return "", nil
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
