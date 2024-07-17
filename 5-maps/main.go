package main

import "fmt"

func main() {

	/*var names map[string]int

	names = make(map[string]int, 0)
	names["cesur"] = 1
	names["bayram"] = 2
	names["serif"] = 3

	fmt.Println(names)
	fmt.Println(names["cesur"])
	fmt.Println(names["fatih"])*/

	names := map[string]int{
		"cesur" :1,
		"bayram" :2,
		"serif":3,
	}
	delete(names, "cesur")
	fmt.Println(names)
}