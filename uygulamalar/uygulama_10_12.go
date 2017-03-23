// uygulama_10_12
package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

// http://sunucu/kullanici/ali şeklinde
// çağırılınca, "ali"ye selam veriyor
func selam(c *gin.Context) {
	isim := c.Param("isim")
	c.String(http.StatusOK, "Merhaba %s", isim)
}

// http://sunucu/kullanici/ali/sil şeklinde
// çağırılınca, "ali" için "sil" eylemini işliyor.
func eylem(c *gin.Context) {
	isim := c.Param("isim")
	eylem := c.Param("eylem")
	mesaj := isim + " için uygulanacak olan eylem: " + eylem
	c.String(http.StatusOK, mesaj)
}
func main() {
	yönlendirici := gin.Default()
	// adres satırından gelen değişkenleri alalım:
	yönlendirici.GET("/kullanici/:isim", selam)
	yönlendirici.GET("/kullanici/:isim/*eylem", eylem)
	yönlendirici.Run(":8000")
}
