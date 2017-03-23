// uygulama_9_14
package main

import (
	"fmt"
	"os"
)

func main() {
	dosya, _ := os.Open("deneme.txt")
	// çıkarken (defer) dosyayı açık unutmayalım:
	defer dosya.Close()
	// Dosya bilgilerini alalım:
	DosyaBilgisi, _ := dosya.Stat()
	fmt.Println("Dosya ismi: ", DosyaBilgisi.Name())
	fmt.Println("Klasör mü: ", DosyaBilgisi.IsDir())
	fmt.Println("Yetkileri: ", DosyaBilgisi.Mode())
	fmt.Println("Değişim tarihi: ", DosyaBilgisi.ModTime())
	fmt.Println("Boyut: ", DosyaBilgisi.Size(), "bayt")
}
