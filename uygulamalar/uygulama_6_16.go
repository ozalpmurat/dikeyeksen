// uygulama_6_16
package main

import "fmt"

func kare(x int) {
	defer func() {
		hata := recover()
		if hata != nil { // recover metni var mı?
			fmt.Println(hata)
			fmt.Println("Panik oluştu. defer kısmı çalıştı.")
		}
	}()
	panic("--- durma noktası ---")
	fmt.Println(x * x)
}
func main() {
	kare(5)
}
