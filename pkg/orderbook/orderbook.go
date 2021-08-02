package orderbook

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// New returns newly initialized OrderBook
func New(targetSize int) *OrderBook {
	return &OrderBook{
		targetSize:   targetSize,
		orders:       map[string]Order{},
		orderIDs:     map[string]OrderSide{},
		totalShares:  map[OrderSide]int{},
		totalHistory: map[TradeAction]float64{},
		canTrade:     map[TradeAction]bool{},
	}
}

// Process is main execution function
func (ob *OrderBook) Process(scanner *bufio.Scanner) error {

	for scanner.Scan() {
		inputString := scanner.Text()
		order, err := ob.parse(inputString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if order.Operation == AddOrder {
			err := ob.addOrder(order)
			if err != nil {
				return err
			}
		} else {
			err := ob.reduceOrder(order)
			if err != nil {
				return err
			}
		}

		err = ob.processOrder(order)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ob *OrderBook) parse(inputString string) (Order, error) {
	in := strings.Split(inputString, " ")

	var order Order
	{
		switch len(in) {
		case 6:
			// format: "timestamp operation id order_type price shares"
			price, _ := strconv.ParseFloat(in[4], 32) // ignoring error for now
			shares, _ := strconv.Atoi(in[5])
			orderSide := getOrderSide(in[3])
			order = Order{
				Timestamp: in[0],
				ID:        in[2],
				Side:      orderSide,
				Operation: AddOrder,
				Size:      shares,
				Price:     floatToFixedSign(price, 2),
			}

		case 4:
			// format: "timestamp operation id shares"
			shares, _ := strconv.Atoi(in[3])

			timestamp := in[0]
			id := in[2]

			order = Order{
				Timestamp: timestamp,
				Operation: ReduceOrder,
				ID:        id,
				Size:      shares,
			}
		default:
			return Order{}, errors.New(fmt.Sprintf("failed to parse order %#q", inputString))
		}
	}

	return order, nil
}

func (ob *OrderBook) addOrder(order Order) error {
	ob.orderIDs[order.ID] = order.Side
	ob.orders[order.ID] = order
	ob.totalShares[order.Side] += order.Size

	return nil
}

func (ob *OrderBook) reduceOrder(order Order) error {
	targetOrder := ob.orders[order.ID]
	targetOrder.Size -= order.Size

	ob.totalShares[targetOrder.Side] -= order.Size

	if targetOrder.Size <= 0 {
		delete(ob.orders, targetOrder.ID)
	} else {
		ob.orders[order.ID] = targetOrder
	}

	return nil
}

func (ob *OrderBook) executeOrder(action TradeAction) float64 {
	if action == Buy {
		return ob.buyShares()
	} else {
		return ob.sellShares()
	}
}

func (ob *OrderBook) buyShares() float64 {
	orders := make([]Order, len(ob.orders))

	for _, order := range ob.orders {
		if order.Side == AskOrder {
			orders = append(orders, order)
		}
	}

	// order by lowest price
	By(sortByPriceAsc).Sort(orders)
	var gainedShares, currentSharesOnBuy int
	var expense float64
	for gainedShares < ob.targetSize {
		for _, order := range orders {
			currentSharesOnBuy = order.Size
			if gainedShares+order.Size >= ob.targetSize {
				currentSharesOnBuy = ob.targetSize - gainedShares
			}

			expense += float64(currentSharesOnBuy) * order.Price
			gainedShares += currentSharesOnBuy

		}
	}

	return expense
}

func (ob *OrderBook) sellShares() float64 {
	var orders []Order
	for _, order := range ob.orders {
		if order.Side == BidOrder {
			orders = append(orders, order)
		}
	}
	// order by highest price
	By(sortByPriceDesc).Sort(orders)
	var gainedShares, currentSharesOnSale int
	var income float64
	for gainedShares < ob.targetSize {
		for _, order := range orders {
			currentSharesOnSale = order.Size
			if gainedShares+order.Size >= ob.targetSize {
				currentSharesOnSale = ob.targetSize - gainedShares
			}

			income += float64(currentSharesOnSale) * order.Price
			gainedShares += currentSharesOnSale

		}
	}

	return income
}

func (ob *OrderBook) processOrder(order Order) error {
	targetOrderSide := ob.orderIDs[order.ID]
	if targetOrderSide == BidOrder {
		if ob.canTrade[Sell] {
			// Enough bids for sale
			if ob.totalShares[BidOrder] >= ob.targetSize {
				total := ob.executeOrder(Sell)
				if total != ob.totalHistory[Sell] {
					ob.totalHistory[Sell] = total
					fmt.Println(formatResult(total, order.Timestamp, "S"))
				}
			} else {
				ob.canTrade[Sell] = false
				fmt.Println(formatResult(0.0, order.Timestamp, "S"))
			}
		} else if !ob.canTrade[Sell] {
			if ob.totalShares[BidOrder] >= ob.targetSize {
				total := ob.executeOrder(Sell)
				ob.totalHistory[Sell] = total
				ob.canTrade[Sell] = true
				fmt.Println(formatResult(total, order.Timestamp, "S"))
			}
		}
	}

	if targetOrderSide == AskOrder {
		if ob.canTrade[Buy] {
			// Enough asks to buy
			if ob.totalShares[AskOrder] >= ob.targetSize {
				total := ob.executeOrder(Buy)
				if total != ob.totalHistory[Buy] {
					ob.totalHistory[Buy] = total
					fmt.Println(formatResult(total, order.Timestamp, "B"))
				}
			} else {
				ob.canTrade[Buy] = false
				fmt.Println(formatResult(0.0, order.Timestamp, "B"))
			}
		} else if !ob.canTrade[Buy] {
			if ob.totalShares[AskOrder] >= ob.targetSize {
				total := ob.executeOrder(Buy)
				ob.totalHistory[Buy] = total
				ob.canTrade[Buy] = true
				fmt.Println(formatResult(total, order.Timestamp, "B"))
			}
		}
	}

	return nil
}

func getOrderSide(str string) OrderSide {
	switch str {
	case "B":
		return BidOrder
	case "S":
		return AskOrder
	default:
		return Undefined
	}
}
