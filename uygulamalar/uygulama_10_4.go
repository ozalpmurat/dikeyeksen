// uygulama_10_4
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	r.GET("/", AnaSayfa)
	r.GET("/gizli", GizemliSeylerYap)
	http.ListenAndServe(":8000", r)
}
func AnaSayfa(Cevap http.ResponseWriter,
	r *http.Request, p httprouter.Params) {
	fmt.Fprint(Cevap, `
<html><body>
<h1>Ana sayfa</h1>
<a href="/gizli">Gizli giriş burada...</a>
</body></html>`)
}
func GizemliSeylerYap(Cevap http.ResponseWriter,
	r *http.Request, p httprouter.Params) {
	fmt.Fprint(Cevap, "Kimseye söyleme!")
}
