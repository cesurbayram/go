package main

import (
	"fmt"
	
)

func main() {

	/* var productName string
	var  quantity int
	var  discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	discount = .37
	isInStock = true
	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock)) */

	/* var productName string = "Golang dersleri"
	fmt.Println(productName,reflect.TypeOf(productName))*/

	/*var quantity int64 = 10
	fmt.Println(quantity, reflect.TypeOf(quantity))*/
     
	var productName string
	var  quantity int
	var  discount float64
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	discount = .37
	isInStock = true

	//fmt.Println("Product name:", productName, "Quantity",quantity, "Discount:", discount, "IsInStock:", isInStock)
	fmt.Printf("Product name: %v, Quantity: %v , Discount: %v, IsInstock %v\n", productName, quantity, discount, isInStock)

}