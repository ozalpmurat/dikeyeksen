// uygulama_6_4
package main

import "fmt"

// Global değişkenler fonksiyon dışında tanımlanır
var a, b int = 1, 2

func değiştir() {
	gecici := a
	a = b
	b = gecici
}
func main() {
	fmt.Println(a, b)
	değiştir()
	fmt.Println(a, b)
}
