package coffee

import "fmt"

// Product represents a type of product available at the coffee shop
type Product int

const (
	Americano Product = iota
	Latte
	FlatWhite
	Filter
)

const (
	// AmericanoPrice is the price of the product Americano
	AmericanoPrice float64 = 2.20
	// LattePrice is the price of the product Latte
	LattePrice float64 = 2.30
	// FlatWhitePrice is the price of the product FlatWhite
	FlatWhitePrice float64 = 2.40
	// FilterPrice is the price of the product Filter
	FilterPrice float64 = 3.50
)

const (
	exactChange string = "Here is your %s, have a nice day!"
	notExact    string = "Sorry, exact change only, try again tomorrow!"
)

// String gives the text representation of the product
func (p Product) String() string {
	switch p {
	case Americano:
		return "Americano"
	case Latte:
		return "Latte"
	case FlatWhite:
		return "Flat White"
	case Filter:
		return "Filter"
	default:
		return ""
	}
}

// Price returns of the price of the product
func (p Product) Price() float64 {
	switch p {
	case Americano:
		return AmericanoPrice
	case Latte:
		return LattePrice
	case FlatWhite:
		return FlatWhitePrice
	case Filter:
		return FilterPrice
	default:
		return float64(0)
	}
}

// Order represents an order submitted to the coffee shop
type Order struct {
	Items  []Product // Items is the collection of products being ordered
	Tender float64   // Tender is the amount of money given for the order
}

// Total calculates the total value of the order from all the products in it
func (o Order) Total() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Price()
	}
	return total
}

func (o Order) String() string {
	count := len(o.Items)

	switch count {
	case 0:
		return ""
	case 1:
		return o.Items[0].String()
	case 2:
		return fmt.Sprintf("%s and %s", o.Items[0].String(), o.Items[1].String())
	default:
		var str string

		for i, item := range o.Items {
			if i+1 == count {
				str += "and "
			}
			str += item.String()
			if i+1 != count {
				str += ", "
			}
		}

		return str
	}
}

// Purchase responds to an order submitted at the coffee shop
func Purchase(order Order) string {
	if order.Total() == order.Tender {
		return fmt.Sprintf(exactChange, order)
	}

	return notExact
}
