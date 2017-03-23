// uygulama_8_11
package main

import "time"
import "fmt"

func uzun_islem(K chan string) {
	time.Sleep(time.Second * 5)
	K <- "uzun"
}
func kisa_islem(K chan string) {
	time.Sleep(time.Second * 2)
	K <- "kısa"
}
func zaman_yaz() { fmt.Println(time.Now()) }
func main() {
	zaman_yaz()
	k1 := make(chan string)
	k2 := make(chan string)
	go uzun_islem(k1)
	go kisa_islem(k2)
	for i := 1; i <= 2; i++ {
		select {
		case mesaj1 := <-k1:
			fmt.Println(mesaj1, "iş tamamlandı")
		case mesaj2 := <-k2:
			fmt.Println(mesaj2, "iş tamamlandı")
		}
	}
	zaman_yaz()
}
