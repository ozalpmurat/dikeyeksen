// uygulama_10_15
package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
	r := gin.Default()
	r.Static("/resimler", "./resimler")
	r.StaticFile("/favicon.ico", "./resimler/favicon.ico")
	r.Run(":8000")
}
