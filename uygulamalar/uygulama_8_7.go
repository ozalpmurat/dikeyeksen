// uygulama_8_7
package main

import "fmt"

func main() {
	sayılar := make(chan string, 5)
	sayılar <- "1"
	sayılar <- "2"
	sayılar <- "3"
	// Kanalı kapatalım:
	close(sayılar)
	for sayı := range sayılar {
		fmt.Println(sayı)
	}
}
