// uygulama_10_1
// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("statik")))
	log.Println("TCP 8000 portunu dinliyorum...")
	http.ListenAndServe(":8000", nil)
}
