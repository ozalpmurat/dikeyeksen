// uygulama_6_19
package main

import "fmt"

func fonk1(ülke *string) {
	// *ülke = "Türkiye"
	*ülke = *ülke + " Cumhuriyeti"
}
func main() {
	metin := "Türkiye"
	fmt.Println(metin) // metin = "Türkiye"
	fonk1(&metin)      // metin'in adresini fonksiyona bildir
	fmt.Println(metin) // metin = "Türkiye Cumhuriyeti"
}
