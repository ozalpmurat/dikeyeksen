// uygulama_9_15
package main

import (
	"fmt"
	"os"
)

func main() {
	dosya, hata := os.Open("deneme.txt")
	if hata != nil { // dosya açılamadıysa:
		fmt.Println("Dosya açılamadı. Gelen hata mesajı:\n", hata)
		return
	}
	// çıkarken (defer) dosyayı açık unutmayalım:
	defer dosya.Close()
	fmt.Println("Dosya başarılı bir şekilde açıldı")
	// buradan sonra dosya ile ilgili
	// işlemleri devam ettirebiliriz.
}
