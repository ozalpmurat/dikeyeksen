// uygulama_6_17
package main

import "fmt"

func kare(işaretçi *int) {
	fmt.Println(*işaretçi * *işaretçi)
}
func main() {
	sayı := 5
	kare(&sayı)
}
