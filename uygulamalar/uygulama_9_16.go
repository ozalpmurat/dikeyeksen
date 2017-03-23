// uygulama_9_16
package main

import (
	"fmt"
	"os"
)

func main() {
	dosya, hata := os.Open("deneme.txt")
	if hata != nil {
		//dosya açılamazsa fonksiyondan çık
		return
	}
	defer dosya.Close()
	// Dosya bilgilerini al:
	DosyaBilgisi, _ := dosya.Stat()
	// Dosya boyutu kadar değişken oluştur:
	BaytKesit := make([]byte, DosyaBilgisi.Size())
	// Dosya içeriğini oku:
	okunanBayt, _ := dosya.Read(BaytKesit)
	// Okunan bayt verisini metin biçimine dönüştür:
	okunanMetin := string(BaytKesit)
	fmt.Println(dosya.Name(),
		"isimli dosyadan okunan",
		okunanBayt, "bayt şu şekildedir:\n",
		okunanMetin)
}
