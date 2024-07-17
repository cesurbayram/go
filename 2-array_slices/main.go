package main

import "fmt"

func main() {
	// fix

	//var names [3]string
	//names [0] ="Cesur"
	//names [1] ="Bayram"
	//names [2] = "Serif"

	/*var names = [4]string{"Cesur","Bayram", "Serif", "fatih"}

	names[1] = "fatih"

	fmt.Println(names[0:2])*/

	var names = []string{"cesur","serif","bayram"}

	names = append(names, "fatih")
	fmt.Println(names)
	
}