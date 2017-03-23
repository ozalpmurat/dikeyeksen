// uygulama_6_14
package main

import "fmt"

func kare(x int) {
	panic("durma noktası") // buraya kadar sağlam mı?
	fmt.Println(x * x)
}
func main() {
	kare(5)
}
