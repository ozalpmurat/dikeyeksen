// uygulama_6_9
package main

import "fmt"

func main() {
	fmt.Println(topla(-5, 4))
	fmt.Println(topla(1, 9, 2, 3))
}
func topla(argumanlar ...int) int {
	var toplam int
	for _, i := range argumanlar {
		toplam += i
	}
	return toplam
}
