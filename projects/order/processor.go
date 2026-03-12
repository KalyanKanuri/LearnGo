package order

import (
	"fmt"
	"math/rand"
)

/* blue print of order processing
> Inventory of items
> cart for items to be added
> order id creation with items in cart
> price calculation based on items in cart
> discount and other calculations based on flags
> return order details
*/

var inventory = map[string]float64{
	"Shirts": 500,
	"Pants":  700,
	"Shoes":  1200,
	"Bags":   1500,
}

var cart []string

type Order map[string]any

func AddToCart(items []string) []string {
	for _, item := range items {
		cart = append(cart, item)
	}
	return cart
}

func genOrderID() string {
	orderID := fmt.Sprintf("ORD-%d", rand.Intn(1000000))
	return orderID
}

func calculatePrice(cart []string, applyDiscount bool) []float64 {
	var totalPrice float64
	var basePrice float64
	var discount float64
	var prices []float64

	for _, item := range cart {
		totalPrice += inventory[item]
		basePrice += inventory[item]
	}
	if applyDiscount {
		totalPrice *= 0.8
		discount = basePrice - totalPrice
	}
	prices = append(prices, basePrice, totalPrice, discount)
	return prices
}

func ProcessOrder(cart []string, applyDiscount bool) Order {
	fmt.Println("-- Processing Order --")
	orderID := genOrderID()
	price := calculatePrice(cart, applyDiscount)
	return Order{
		"orderID":    orderID,
		"items":      cart,
		"basePrice":  price[0],
		"finalPrice": price[1],
		"discount":   price[2],
	}
}
