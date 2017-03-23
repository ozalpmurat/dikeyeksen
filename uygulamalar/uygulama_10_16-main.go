// uygulama_10_16
// main.go
package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

func AnaSayfa(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		// Şablona gönderilecek olan değişkenler:
		"başlık": "Go programcıları",
		"mesaj":  "Herkese keyifli kodlamalar dileriz",
	})
}
func main() {
	r := gin.Default()
	// Şablon dosyalarını yükle:
	r.LoadHTMLGlob("templates/*")
	r.GET("/", AnaSayfa)
	r.Run(":8000")
}
