// uygulama_5_3
package main

import "fmt"

func main() {
	aday := [5]int{
		81,
		100,
		27,
		95,
		45,
	}
	var adaysayisi int = len(aday)
	for i := 0; i < adaysayisi; i++ {
		if aday[i] < 50 {
			fmt.Println(i, ". aday başarısız oldu:", aday[i])
		}
	}
}
