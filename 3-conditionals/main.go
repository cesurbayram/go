package main

import "fmt"

func main() {
	/*var age = 17
	if age >= 18 {
		fmt.Println("You can vote!")
	}else{
		fmt.Println(" You can not vote!")
	}*/
	a := 10
	b := 20
	c := 30

		if a >=b && a >=c{
			fmt.Println("Greatest variable is a!")
		}else if b>=a && b>=c{
			fmt.Println("Greatest variables is b!")
		}else{
			fmt.Println("Greatest variables is c!")
		}
}