// uygulama_8_2
package main

import "fmt"
import "time"

func HarfYaz(i int) {
	harfler := "abcçd"
	for _, harf := range harfler {
		fmt.Printf("%d%s ", i, string(harf))
	}
	fmt.Println()
}
func main() {
	// HarfYaz fonksiyonunu 5 kere sıralı olarak çağıralım.
	fmt.Println("--- sıralı çalışan ---")
	for i := 0; i < 5; i++ {
		HarfYaz(i) // fonksiyon çağır
	}
	// HarfYaz fonksiyonunu 5 kere eşzamanlı olarak çağıralım.
	fmt.Println("--- eşzamanlı çalışan ---")
	for i := 0; i < 5; i++ {
		go HarfYaz(i) // gorutin olarak çağır
	}
	// 1 saniye bekle, fonksiyonlar tamamlansın.
	time.Sleep(time.Millisecond * 100)
}
