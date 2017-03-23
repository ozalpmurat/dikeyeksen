// uygulama_9_22
package main

import (
	"io"
	"net/http"
)

// Sunucunun ana sayfası =>> http://sunucu_adresi/
func AnaSayfa(cevap http.ResponseWriter, istek *http.Request) {
	cevap.Header().Set("Content-Type", "text/html")
	io.WriteString(cevap, `
<DOCTYPE html>
<html>
<body>
<h1>Ana sayfaya geldiniz</h1>
<a href="/alt_klasor"> ikinci sayfa icin tiklayin</a>
</body>
</html>`)
}

// ikinci sayfa =>> http://sunucu_adresi/merhaba
func ikinci_sayfa(cevap http.ResponseWriter, istek *http.Request) {
	cevap.Header().Set("Content-Type", "text/html")
	io.WriteString(cevap, "Go web sunucusundan merhaba!")
}
func main() {
	//Gelen istekteki adrese göre fonksiyona yönlendir:
	http.HandleFunc("/", AnaSayfa)
	http.HandleFunc("/alt_klasor/", ikinci_sayfa)
	http.ListenAndServe(":1983", nil)
}
