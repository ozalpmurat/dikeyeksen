// uygulama_5_5
package main

import "fmt"

func main() {
	aday := [5]int{81, 100, 27, 95, 45}
	basarili := 0
	for _, puan := range aday {
		if puan >= 50 {
			basarili += 1
		}
	}
	fmt.Println("Başarılı aday sayısı:", basarili)
}
