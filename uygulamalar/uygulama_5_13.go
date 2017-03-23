// uygulama_5_13
package main

import "fmt"

func main() {
	il_nufus := map[string]int{
		"Bilecik":  210000,
		"İstanbul": 15000000,
		"Sakarya":  1000000,
	}
	for anahtar, deger := range il_nufus {
		fmt.Printf("İl: %s, Nüfus: %d\n", anahtar, deger)
	}
}
