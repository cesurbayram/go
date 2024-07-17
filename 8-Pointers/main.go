package main

import "fmt"

func main() {
	/*var a int
	var p *int
	a = 10;
	p = &a
	*p = 20;
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(a) */

	/*var a = 10
	var b int 
	var p *int 
	b = a
	p = &a
	*p = 20
	fmt.Println(a, b)*/

	// var a int = 10

	/*fmt.Println(a)
	add12pointer(&a)
	fmt.Println(a)*/

	numbers := []int{1, 2, 3}
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)
}

func changeValue(numbers []int) {
	numbers[0] = 1000
}

func add12(x int) { // değer ile geçme
	x = x + 12
}

func add12pointer(x *int) { // referans ile geçme
	*x = *x + 12
}
