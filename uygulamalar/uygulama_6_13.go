// uygulama_6_13
package main

import "fmt"

func main() {
	işlem_yap()
}
func işlem_yap() {
	i := 0
	defer fmt.Println(i) // 2. sırada işletilecek
	i++
	defer fmt.Println(i) // 1. sırada işletilecek
	fmt.Println("Bu satır defer'lerden önce işletilir.")
}
