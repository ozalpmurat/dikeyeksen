// uygulama_4_5
package main

import "fmt"

func main() {
	var i, b2, b3, b4 int
	for i = 1; i <= 100; i++ {
		if i%2 == 0 {
			// 2'ye bölünebilen
			b2++
		} else if i%3 == 0 {
			// 3'e bölünebilen
			b3++
		} else if i%4 == 0 {
			// 4'e bölünebilen
			b4++
		}
	}
	fmt.Printf("2:%d, 3:%d, 4:%d", b2, b3, b4)
}
