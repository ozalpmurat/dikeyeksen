// uygulama_3_4
package main

import (
	"fmt"
)

func main() {
	var isim, parola, eposta bool
	isim = true
	parola = true
	eposta = false
	fmt.Println(isim && parola)
	fmt.Println(eposta || isim)
	fmt.Println(!isim)
	fmt.Println("Sakarya" == "Adapazarı")
}
