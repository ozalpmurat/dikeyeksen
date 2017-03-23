// uygulama_9_2
package main

import (
	"fmt"
	"strings"
)

func main() {
	adresler := []string{
		"ali@example.com",
		"veli.example.com",
		"ayşe@com"}
	for _, adres := range adresler {
		if !strings.ContainsAny(adres, "@") {
			fmt.Println("Hatalı e-posta adresi: ", adres)
		}
	}
}
