// uygulama_5_2
package main

import "fmt"

func main() {
	var aday [5]float32
	var toplam float32
	aday[0] = 81
	aday[1] = 100
	aday[2] = 65
	aday[3] = 95
	aday[4] = 45
	for i := 0; i <= 4; i++ {
		toplam += aday[i]
	}
	fmt.Println("Toplam aday sayısı=", len(aday))
	fmt.Println("Ortalama=", toplam/5)
}
