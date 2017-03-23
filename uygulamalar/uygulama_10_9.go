// uygulama_10_9
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Kitap struct { // sorgu sonucunu aktaracağımız yapı
	isim  string
	yazar string
}

func main() {
	vt, _ := sql.Open("sqlite3", "kitaplar.sqlite")
	listele(vt)
}
func listele(vt *sql.DB) {
	// Sorguyu çalıştır
	Satirlar, _ := vt.Query("SELECT * FROM kitap")
	defer Satirlar.Close()
	// Yeni bir kesit oluştur
	Kitaplar := make([]*Kitap, 0)
	// SQL cevabının her satırı için dön
	for Satirlar.Next() {
		satir := new(Kitap)
		// satir yapısının alanlarını doldur
		Satirlar.Scan(&satir.isim, &satir.yazar)
		// yeni satır alanlarını Kitaplar kesitine ekle
		Kitaplar = append(Kitaplar, satir)
	}
	for i, kitap := range Kitaplar {
		fmt.Printf("%d) %s, %s\n", i, kitap.yazar, kitap.isim)
	}
}
