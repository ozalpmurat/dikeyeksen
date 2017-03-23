// uygulama_8_6
package main

import "fmt"

func main() {
	tamponluKanal := make(chan string, 2)
	tamponluKanal <- "birinci mesaj"
	tamponluKanal <- "ikinci mesaj"
	// Kanaldaki mesajları alalım
	fmt.Println(<-tamponluKanal)
	fmt.Println(<-tamponluKanal)
}
