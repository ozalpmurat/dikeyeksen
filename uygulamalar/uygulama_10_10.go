// uygulama_10_10
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type Kitap struct {
	isim  string
	yazar string
}

// VT değişkenini global olarak tanımladığımız
// için fonksiyona argüman göndermeye gerek yok:
var VT *sql.DB

func main() {
	VT, _ = sql.Open("sqlite3", "kitaplar.sqlite")
	http.HandleFunc("/", listele)
	http.ListenAndServe(":8000", nil)
}
func listele(HTMLcevap http.ResponseWriter,
	HTMListek *http.Request) {
	// sorgu cevabı kesite aktarılsın:
	Satirlar, _ := VT.Query("SELECT * FROM kitap")
	defer Satirlar.Close()
	Kitaplar := make([]*Kitap, 0)
	for Satirlar.Next() {
		satir := new(Kitap)
		Satirlar.Scan(&satir.isim, &satir.yazar)
		Kitaplar = append(Kitaplar, satir)
	}
	// HTML cevabını oluşturalım:
	for i, kitap := range Kitaplar {
		fmt.Fprintf(HTMLcevap,
			"%d) %s - %s\n",
			i, kitap.yazar, kitap.isim)
	}
}
