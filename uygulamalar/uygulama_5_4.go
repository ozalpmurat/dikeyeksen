// uygulama_5_4
package main

import "fmt"

func main() {
	aday := [5]int{81, 100, 27, 95, 45}
	for i, puan := range aday {
		if puan < 50 {
			fmt.Println(i, ". aday başarısız oldu:", puan)
		}
	}
}
