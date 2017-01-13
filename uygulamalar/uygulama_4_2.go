// uygulama_4_2
package main

import "fmt"

func main() {
	i := 1
	toplam := 0
	for i <= 5 { // başlangıç ifadesi ve ilerleme kuralı pas geçilebilir.
		i++
		toplam = toplam + i
	}
	fmt.Println("Toplam:", toplam)
}
