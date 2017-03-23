// uygulama_5_15
package main

import "fmt"

func main() {
	var toplam int
	bilgisayar := map[string]int{
		"Dizüstü":  27,
		"Masaüstü": 123,
		"Tablet":   8,
		"Sunucu":   3,
	}
	for _, deger := range bilgisayar {
		var toplam int
		toplam = toplam + deger
	}
	fmt.Printf("Toplam %d bilgisayar var.\n", toplam)
}
