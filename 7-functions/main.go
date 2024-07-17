package main

import "fmt"

func main() {
	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum(3, 4, 5, 6))
}

func sum(numbers ...int) int {
	total := 0 // toplamı tutacak değişken
	for _, value := range numbers {
		total += value
	}
	return total
}

func calculation(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func add(x int, y int) int {
	return x + y
}
