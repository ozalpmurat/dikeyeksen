// uygulama_8_14
package main

import (
	"fmt"
	"time"
)

func main() {
	sayaç := time.NewTimer(time.Second * 3)
	fmt.Println("3 saniye geri sayım başlatıldı...")
	go func() {
		<-sayaç.C
		fmt.Println("Sayaç tamamlandı")
	}()
	// İstenildiğinde timer durdurulabilir:
	time.Sleep(time.Second * 1)
	iptal := sayaç.Stop()
	fmt.Println("Sayaç durduruldu mu? ", iptal)
}
