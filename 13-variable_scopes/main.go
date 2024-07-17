package main

import "fmt"

func main() {
	var x int  // x'i burada tanımlıyoruz, böylece diğer if bloğu dışında da erişilebilir
	var condition = true

	if condition {
		x = 10 // x'i burada atıyoruz
		fmt.Println(x)
	}
	fmt.Println(x) // if bloğu dışında x'e erişebiliriz

	condition = true // condition değişkenini burada tekrar atıyoruz
	if condition {
		fmt.Println(condition)
	}

	fmt.Println(condition)

}