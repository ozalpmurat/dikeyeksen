// uygulama_10_8
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// VT bağlantısını yap:
	vt, _ := sql.Open("sqlite3", "kitaplar.sqlite")
	// tablo oluştur:
	vt.Exec(`CREATE TABLE IF NOT EXISTS
kitap(isim text, yazar text);`)
	// tabloya 3 tane kitap kaydet:
	vt.Exec(`INSERT INTO kitap VALUES
('Atlas Silkindi','Ayn Rand'),
('Suç ve Ceza','Dostoyevski'),
('Otomatik Portakal','Anthony Burgess');`)
	var isim, yazar string
	// tablonun ilk satırını getir:
	vt.QueryRow("SELECT * FROM kitap;").Scan(&isim, &yazar)
	fmt.Printf("İlk satırdaki kitap: %s - %s\n", yazar, isim)
}
