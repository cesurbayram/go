package main

import "fmt"

func main() {
	//var numbers = []int{1, 2, 3, 4}

	/*for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}*/

	/*for _, value :=range numbers{
		fmt.Println(value)
	}*/

	/*var language = "Golang"

	for _, character := range language {
		fmt.Printf(" Chracter %c\n", character)
	}*/

	var names = map[string]int{
		"Cesur": 10,
		"Bayram": 20,
		"Serif": 30,
	}
	for key, value :=range names {
		fmt.Printf("key: %s , value: %d\n", key, value)
	}

}
