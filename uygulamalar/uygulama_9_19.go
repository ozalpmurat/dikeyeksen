// uygulama_9_19
package main

import "os"

func main() {
	dosya, hata := os.Create("test.txt")
	if hata != nil {
		// dosya oluşturulamadı
		return
	}
	defer dosya.Close()
	dosya.WriteString(`Go, Google'dan bir ekip ile
açık kaynak topluluğundan çok sayıda katılımcı
tarafından geliştirilen açık kaynaklı bir projedir.`)
}
