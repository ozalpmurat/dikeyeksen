// uygulama_5_7
package main

import "fmt"

func main() {
	slice := []int{1, 0, 7, 1}
	for i, değer := range slice {
		fmt.Printf("İndeks: %d Değer: %d\n", i, değer)
	}
}
