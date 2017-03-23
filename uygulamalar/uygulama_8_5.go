// uygulama_8_5
package main

import (
	"fmt"
	"time"
)

func bekle() {
	var enter_tusu string
	fmt.Scanln(&enter_tusu)
	fmt.Println("Program tamamlandı")
}
func sensörden_kanala(K chan string) {
	for {
		K <- "zaman: " + time.Now().String()
	}
}
func kanaldan_ekrana(K chan string) {
	for {
		mesaj := <-K
		fmt.Println(mesaj)
		time.Sleep(time.Second * 2)
	}
}
func main() {
	sensör_veri_yolu := make(chan string)
	fmt.Println("Çıkmak için Enter tuşuna basınız.")
	go kanaldan_ekrana(sensör_veri_yolu)
	go sensörden_kanala(sensör_veri_yolu)
	bekle()
}
