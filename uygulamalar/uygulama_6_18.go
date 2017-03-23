// uygulama_6_18
package main

import "fmt"

func main() {
	sayı := 54
	fmt.Println("Sayı: ", sayı)
	fmt.Println("Sayı'nın adresi: ", &sayı)
	adres := &sayı
	fmt.Println("Adres:", adres)
	fmt.Println("Adres'teki veri:", *adres)
}
