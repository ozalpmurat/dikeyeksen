// uygulama_8_10
package main

import "fmt"

func hoparlor(S <-chan string, ses_seviyesi int) {
	fmt.Printf("Anons ses seviyesi: %d\n", ses_seviyesi)
	fmt.Printf("Anons metni: %s\n", <-S)
}
func mikrofon(S chan<- string, mesaj string) {
	S <- mesaj
}
func main() {
	ses := make(chan string, 1)
	mikrofon(ses, "Cihaz kontrol!")
	hoparlor(ses, 3)
}
