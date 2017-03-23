// uygulama_10_13
package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

func blog(c *gin.Context) {
	//varsayılan yazı numarası 1 olsun:
	yazi := c.DefaultQuery("yazi", "1 (varsayılan)")
	eylem := c.Query("eylem")
	c.String(http.StatusOK,
		"%s numaralı yazı için %s işlemi yapılacaktır.",
		yazi, eylem)
}
func main() {
	yönlendirici := gin.Default()
	yönlendirici.GET("/blog", blog)
	yönlendirici.Run(":8000")
}
