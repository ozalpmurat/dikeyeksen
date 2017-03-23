// uygulama_6_21
package main

import "fmt"

func fonk(x *int) int {
	*x++
	return *x + 1
}
func main() {
	sayı := 0
	fmt.Println(fonk(&sayı))
}
