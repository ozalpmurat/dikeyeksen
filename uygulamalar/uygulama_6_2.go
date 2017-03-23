// uygulama_6_2
package main

import "fmt"

func main() {
	a := 5
	b := 2
	sonuc := topla(a, b)
	fmt.Println("Toplam (fonksiyon cevabı)=", sonuc)
}

// topla fonksiyonu, iki tamsayının toplamını döndürür.
func topla(a, b int) int {
	return a + b
}
