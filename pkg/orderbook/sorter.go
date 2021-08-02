package orderbook

import "sort"

// By is the type of a "less" function that defines the ordering of its Order arguments.
type By func(o1, o2 *Order) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(orders []Order) {
	os := &orderSorter{
		orders: orders,
		by:     by,
	}
	sort.Sort(os)
}

// orderSorter joins a By function and a slice of Orders to be sorted.
type orderSorter struct {
	orders []Order
	by     func(p1, p2 *Order) bool
}

// Len is part of sort.Interface.
func (s *orderSorter) Len() int {
	return len(s.orders)
}

// Swap is part of sort.Interface.
func (s *orderSorter) Swap(i, j int) {
	s.orders[i], s.orders[j] = s.orders[j], s.orders[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *orderSorter) Less(i, j int) bool {
	return s.by(&s.orders[i], &s.orders[j])
}

// sortByPriceAsc is a sort function used to get lower prices first
func sortByPriceAsc(o1, o2 *Order) bool {
	if floatEqual(o1.Price, o2.Price) {
		return o2.Size > o1.Size
	}
	return o1.Price < o2.Price
}

// sortByPrice is a sort function used to get higher prices first
func sortByPriceDesc(o1, o2 *Order) bool {
	if floatEqual(o1.Price, o2.Price) {
		return o2.Size > o1.Size
	}

	return o1.Price >= o2.Price
}
