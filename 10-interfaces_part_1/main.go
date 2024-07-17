package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

type Flower struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}

func (flower *Flower) ShippingCost() int {
	return 12 + flower.desi*3
}

func main() {
	products := []IShippable{
		&Book{desi: 10},
		&Book{desi: 20},
		&Electronic{desi: 20},
		&Flower{desi: 10},
	}

	fmt.Println(calculateTotalShippingCostOfBooks(products))
}

func calculateTotalShippingCostOfBooks(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}
