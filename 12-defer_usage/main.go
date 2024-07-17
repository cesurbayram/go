package main

import "fmt"

func main() {
	
	//ornek 1
	defer fmt.Println("1.Defer")
	defer fmt.Println("2.Defer")
	defer fmt.Println("3.Defer")
	fmt.Println("Main Fonksiyonu")

	//ornek2
	var condition = true
	defer fmt.Println("1.Defer")

	if condition {
		return
	}
	defer fmt.Println("2.Defer")

	//ornek3
	x :=10
	y :=20

	defer fmt.Println("x :", x)
	defer fmt.Println("y :", y)

	x = 100
	y = 200

	fmt.Println("x :", x)
	fmt.Println("y :", y)
}
