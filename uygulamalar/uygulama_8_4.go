// uygulama_8_4
package main

import "fmt"

func müzevir(dedikodu chan string) {
	dedikodu <- "Leyla, Mecnun'u seviyormuş"
}
func main() {
	dedikodu := make(chan string)
	go müzevir(dedikodu) // gorutin çalıştırma
	MahalleGastesi := <-dedikodu
	fmt.Println(MahalleGastesi)
}
