// uygulama_6_6
package main

import "fmt"

func main() {
	var a, b int = 1, 2
	fmt.Println(a, b)
	a, b = b, a // değişken değerlerini değiştir
	fmt.Println(a, b)
}
