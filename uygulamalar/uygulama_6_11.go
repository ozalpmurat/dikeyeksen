// uygulama_6_11
package main

import "fmt"

func main() {
	fmt.Println(faktöryel(4))
}
func faktöryel(sayı int) int {
	çarpım := 1
	for i := 1; i <= sayı; i++ {
		çarpım = çarpım * i
	}
	return çarpım
}
