package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.WaitGroup nesnesi oluştur
	var wg sync.WaitGroup

	// WaitGroup için goroutine'lar için bekleyici sayısını ayarla
	wg.Add(2)

	// İlk goroutine
	go func() {
		defer wg.Done() // goroutine'un tamamlandığını belirt
		fmt.Println("f1")
	}()

	// İkinci goroutine
	go func() {
		defer wg.Done() // goroutine'un tamamlandığını belirt
		fmt.Println("f2")
	}()

	// Bekleme grubunun tamamının tamamlanmasını bekle
	wg.Wait()

	// Ana fonksiyonun sonunu belirt
	fmt.Println("End of the main")
}
