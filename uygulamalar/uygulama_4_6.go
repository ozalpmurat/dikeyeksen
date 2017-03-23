// uygulama_4_6
package main

import (
	"fmt"
)

func main() {
	sayi := 0
	switch {
	case sayi > 0:
		fmt.Println("pozitif")
	default:
		fmt.Println("negatif")
	}
}
