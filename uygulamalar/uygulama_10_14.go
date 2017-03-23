// uygulama_10_14
package main

import "gopkg.in/gin-gonic/gin.v1"
import "fmt"

func blog_post(c *gin.Context) {
	// Query => URL'den gelenler
	no := c.Query("no")
	sayfa := c.Query("sayfa")
	// PostForm => POST verisinden gelenler
	isim := c.PostForm("isim")
	mesaj := c.PostForm("mesaj")
	fmt.Printf("no:%s; sayfa:%s; isim:%s; mesaj:%s\n",
		no, sayfa, isim, mesaj)
}
func main() {
	yonlendirici := gin.Default()
	yonlendirici.POST("/blog_post", blog_post)
	yonlendirici.Run(":8000")
}
