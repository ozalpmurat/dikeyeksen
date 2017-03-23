// uygulama_10_5
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Selamla(w http.ResponseWriter,
	r *http.Request, Arguman httprouter.Params) {
	//argüman olarak gelen ismi alalım
	isim := Arguman.ByName("degisken")
	fmt.Fprintf(w, "Merhaba %s!\n", isim)
}
func main() {
	r := httprouter.New()
	r.GET("/selamla/:degisken", Selamla)
	http.ListenAndServe(":8000", r)
}
