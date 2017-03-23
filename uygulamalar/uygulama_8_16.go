// uygulama_8_16
package main

import "fmt"
import "time"

func işçi(işçi_no int, iş_kuyruğu chan int, iş_cevapları chan int) {
	for iş := range iş_kuyruğu {
		fmt.Println(işçi_no, ". işçi, ", iş, ". işi yapıyor")
		// iş 1 saniye sürsün:
		time.Sleep(time.Second)
		iş_cevapları <- iş
	}
}
func main() {
	iş_sayısı := 9
	işçi_sayısı := 3
	// iş kuyruğu ve iş cevapları için
	// iş sayısı kadar tamponlu kanal oluşturalım
	iş_kuyruğu := make(chan int, iş_sayısı)
	iş_cevapları := make(chan int, iş_sayısı)
	// 3 tane işçi oluşturalım
	for i := 1; i <= işçi_sayısı; i++ {
		go işçi(i, iş_kuyruğu, iş_cevapları)
	}
	// 9 tane iş uydurup, kuyruğa koyalım
	for iş := 1; iş <= iş_sayısı; iş++ {
		iş_kuyruğu <- iş
	}
	// İşlerin cevaplarını bekleyelim
	for a := 1; a <= iş_sayısı; a++ {
		<-iş_cevapları
	}
	fmt.Println("---\nTüm işler tamamlandı.")
}
