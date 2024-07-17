package main

import (
	"fmt"
)

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct{}

func (ps *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{Text: "Ürün ismi boş olamaz", Code: "1001"}
	}
	if product.price < 0 {
		return ValidationError{Text: "Ürün fiyati 0'dan büyük olmali", Code: "1002"}
	}
	fmt.Println("Ürün eklendi!")
	return nil
}

type ValidationError struct {
	Code string
	Text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata: %s, Kod: %s", validationError.Text, validationError.Code)
}

func main() {
	productService := ProductService{}

	err := productService.Add(Product{id: 1, name: "", price: 3000})
	if err != nil {
		// Hata ValidationError türünden mi kontrol edilir.
		// Eğer hata ValidationError ise özel durumları işle
		if validationError, ok := err.(ValidationError); ok {
			fmt.Println("Özel durumlar işleniyor...")
			fmt.Println("Hata Kodu:", validationError.Code)
			fmt.Println("Hata Mesaji:", validationError.Text)
		} else {
			// Aksi takdirde, genel bir hata varsa bu hata mesajı yazdırılır.
			fmt.Println(err)
		}
	}

	err = productService.Add(Product{id: 2, name: "Kalem", price: 10})
	if err != nil {
		fmt.Println(err)
	}
}
