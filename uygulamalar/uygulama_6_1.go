// uygulama_6_1
package main

import "fmt"

func main() {
	a := 5
	b := 4
	topla(a, b)
}

// topla isimli fonksiyon aşağıda tanımlanıyor.
func topla(m, n int) {
	fmt.Println("Sayıların toplamı=", m+n)
}
