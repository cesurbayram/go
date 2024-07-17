package main

import "fmt"

// IEncoder, encode ve decode metodlarını içeren bir arayüzü tanımlar.
type IEncoder interface {
	Encode(value string) // Encode metodunu tanımlar.
	Decode(value string) // Decode metodunu tanımlar.
}

// XEncoder, IEncoder arayüzünü uygulayan bir yapıdır.
type XEncoder struct{} 

// YEncoder, IEncoder arayüzünü uygulayan bir yapıdır.
type YEncoder struct{}

// Encode metodunu XEncoder için uygular.
func (xEncoder *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi.")
}

// Decode metodunu XEncoder için uygular.
func (xEncoder *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi.")
}

// Encode metodunu YEncoder için uygular.
func (yEncoder *YEncoder) Encode(value string) {
	fmt.Println("YEncoder ile encode edildi.")
}

// Decode metodunu YEncoder için uygular.
func (yEncoder *YEncoder) Decode(value string) {
	fmt.Println("YEncoder ile decode edildi.")
}

func main() {
	var encoder IEncoder // IEncoder tipinde bir değişken tanımlanır.
	encoder = &YEncoder{} // YEncoder tipinden bir nesne oluşturulur ve encoder değişkenine atanır.
	encoder.Encode("123456") // encoder değişkeninin Encode metodu çağrılır.
	encoder.Decode("xsad2323sads") // encoder değişkeninin Decode metodu çağrılır.
}
