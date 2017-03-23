// uygulama_10_3
// main.go
package main

import (
	"html/template"
	"net/http"
)

type Söz struct {
	Söyleyen string
	Zaman    int
	Cümle    string
}

func SözListele(HTMLcevap http.ResponseWriter,
	HTMListek *http.Request) {
	var ÖzlüSözler [5]Söz
	ÖzlüSözler[0] = Söz{"Thomas Watson (IBM)", 1943,
		"Dünyada belki de en fazla 5 bilgisayara ihtiyaç vardır."}
	ÖzlüSözler[1] = Söz{"Ken Olson (DEC)", 1977,
		`Bence hiç kimsenin evine bilgisayar sokmak için
geçerli nedeni olamaz.`}
	ÖzlüSözler[2] = Söz{"Darryl Zanuck (film yapımcısı)", 1946,
		`Televizyonun modası geçecektir.
Kim her akşam aynı aptal kutuya bakmak ister ki ?`}
	ÖzlüSözler[3] = Söz{"Murat Özalp (this->Author)", 1993,
		`Tarayıcıdan bilgisayara aktarılan yazılar
asla metin olarak işlenemeyecektir.`}
	şablon, _ := template.ParseFiles("templates/index.html")
	şablon.Execute(HTMLcevap, ÖzlüSözler)
}
func main() {
	http.HandleFunc("/", SözListele)
	http.ListenAndServe(":8000", nil)
}
