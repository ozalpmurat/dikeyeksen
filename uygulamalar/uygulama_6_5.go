// uygulama_6_5
package main

import "fmt"

func değiştir(x, y int) (int, int) {
	return y, x
}
func main() {
	// Yerel değişkenler fonksiyon içinde tanımlanır
	a := 1
	b := 2
	fmt.Println(a, b)
	a, b = değiştir(a, b)
	fmt.Println(a, b)
}
