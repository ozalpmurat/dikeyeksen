// uygulama_8_9
package main

import "fmt"

func main() {
	kontrol := make(chan bool)
	go func() {
		println("gorutin'den gelen mesaj")
		close(kontrol)
	}()
	fmt.Println("main'den gelen mesaj")
	<-kontrol
}
