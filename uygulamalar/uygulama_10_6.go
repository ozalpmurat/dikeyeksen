// uygulama_10_6
package main

import (
	"encoding/json"
	"net/http"
)

type Söz struct {
	Söyleyen string
	Zaman    int
	Cümle    string
}

func SözListele(cevap http.ResponseWriter, istek *http.Request) {
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
	// JSON cevabını oluştur:
	liste, _ := json.Marshal(ÖzlüSözler)
	// JSON özelliklerini ayarlarla:
	cevap.Header().Set("Content-Type",
		"application/json; charset=utf-8")
	// JSON cevabını gönder:
	cevap.Write(liste)
}
func main() {
	http.HandleFunc("/", SözListele)
	http.ListenAndServe(":8000", nil)
}
