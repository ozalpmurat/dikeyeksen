// uygulama_4_4
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // 2'ye bölümden kalan 0 mı?
			// sayı çift ise çalışsın
			fmt.Printf("%d ", i)
		}
	}
}
