package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "Cesur Bayram Serif", age: 30, 
		address: Address{city: "Istanbul", district: "Atasehir"}}

	customer1.ToString()
	customer1.changeName("Cesur Bayram")
	customer1.ToString()
}
func (customer *Customer) changeName(newName string){
	customer.name = newName
}
type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}

func (customer Customer) ToString() {
	fmt.Printf("Name: %s, Age: %d\n", customer.name, customer.age)
}
