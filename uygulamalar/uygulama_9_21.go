// uygulama_9_21
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	// example.org sitesine bağlan, http cevabını al:
	http_cevap, _ := http.Get("http://example.org/")
	// gelen http cevabının gövdesini sunucudan oku:
	html_sayfa, _ := ioutil.ReadAll(http_cevap.Body)
	http_cevap.Body.Close()
	// html web sayfasının kodlarını ekrana dök:
	fmt.Printf("%s\n------\n", html_sayfa)
	fmt.Println(reflect.TypeOf(http_cevap))
	fmt.Println(reflect.TypeOf(html_sayfa))
}
