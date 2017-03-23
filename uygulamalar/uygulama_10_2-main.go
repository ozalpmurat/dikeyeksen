// uygulama_10_2
// main.go
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Söz struct {
	Cümle    string
	Söyleyen string
}

func SözListele(HTMLcevap http.ResponseWriter,
	HTMListek *http.Request) {
	ÖrnekSöz1 := Söz{
		"Her şeyi olabildiğince sade yapın, ama basit değil.",
		"Albert Einstein"}
	şablon, _ := template.ParseFiles("templates/index.html")
	şablon.Execute(HTMLcevap, ÖrnekSöz1)
	log.Println(ÖrnekSöz1.Söyleyen, ": ", ÖrnekSöz1.Cümle)
}
func main() {
	http.HandleFunc("/", SözListele)
	http.ListenAndServe(":8000", nil)
}
