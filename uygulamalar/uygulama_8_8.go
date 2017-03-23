// uygulama_8_8
package main

import (
	"fmt"
	"time"
)

func formatla(kontrol_yolu chan string) {
	fmt.Println("Sabit disk biçimlendiriliyor...")
	time.Sleep(time.Second * 2)
	kontrol_yolu <- "Temizlik tamamlandı."
}
func main() {
	kontrol_yolu := make(chan string, 1)
	go formatla(kontrol_yolu)
	<-kontrol_yolu
}
