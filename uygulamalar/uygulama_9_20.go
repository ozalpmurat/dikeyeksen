// uygulama_9_20
package main

import (
	"fmt"
	"os"
)

func listele(yol string) (int, int64) {
	klasör, _ := os.Open(yol)
	defer klasör.Close()
	// Readdir(x) x<0 ise "tüm dosyalar" demek.
	dosyalar, _ := klasör.Readdir(-1)
	fmt.Println("Boyut\tDosya adı\n", "-----\t--------")
	// Dosya boyut toplamı için int64 kullanalım:
	var toplam int64
	for _, dosya := range dosyalar {
		fmt.Print(dosya.Size(), "\t")
		fmt.Println(dosya.Name())
		toplam = toplam + dosya.Size()
	}
	// Dosya sayısı ve toplam boyutu döndürelim:
	return len(dosyalar), toplam
}
func main() {
	var dosya_sayisi int
	var klasor_boyutu int64
	dosya_sayisi, klasor_boyutu = listele("/home/murat/Download/")
	fmt.Printf("Toplam %d dosya, %dKB yer kaplıyor.\n",
		dosya_sayisi, klasor_boyutu/1024)
}
