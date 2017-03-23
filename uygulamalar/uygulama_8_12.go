// uygulama_8_12
package main

import "fmt"
import "time"

func işyap(K chan bool, saniye time.Duration) {
	// iş yapıyorum
	time.Sleep(saniye * time.Second)
	K <- true
}
func süretut(K chan bool, saniye time.Duration) {
	// bekliyorum
	time.Sleep(saniye * time.Second)
	K <- true
}
func main() {
	gerisayım := make(chan bool, 1)
	işcevabı := make(chan bool, 1)
	//kaç saniyelik iş yapsın
	go işyap(işcevabı, 3)
	// İşin bitmesini kaç saniye bekleyelim
	go süretut(gerisayım, 2)
	select {
	case <-işcevabı:
		fmt.Println("süre bitmeden iş yetişti")
	case <-gerisayım:
		fmt.Println("zaman aşımı oldu")
	}
}
