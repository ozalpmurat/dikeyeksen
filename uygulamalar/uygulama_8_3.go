// uygulama_8_3
package main

import "fmt"

func main() {
	kanal1 := make(chan string)         //1. satır
	go func() { kanal1 <- "merhaba" }() //2. satır
	mesaj := <-kanal1                   //3. satır
	fmt.Println(mesaj)
}
