// uygulama_10_11
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Urun struct {
	gorm.Model
	Kod   string
	Fiyat uint
}

func Listele(vt *gorm.DB) {
	// gelen kayıtlar kesit türüne yerleştirilsin
	var urun []Urun
	// SELECT * FROM uruns;
	vt.Find(&urun)
	fmt.Println("\n------ Güncel Ürün Listesi ------")
	for _, kayit := range urun {
		fmt.Println(kayit.ID, "\t",
			kayit.Kod, "\t", kayit.Fiyat)
	}
}
func UrunEkle(vt *gorm.DB) {
	// INSERT INTO uruns VALUES ('i1453',1000),
	// ('m1070',2000),('t1923',3000);
	vt.Create(&Urun{Kod: "i1453", Fiyat: 1000})
	vt.Create(&Urun{Kod: "m1070", Fiyat: 2000})
	vt.Create(&Urun{Kod: "t1923", Fiyat: 3000})
	fmt.Println("*** ürünler eklendi...")
}
func main() {
	vt, _ := gorm.Open("sqlite3", "stok.db")
	defer vt.Close()
	vt.AutoMigrate(&Urun{}) // Tabloyu oluştur
	UrunEkle(vt)
	Listele(vt)
	// SELECT fiyat FROM uruns WHERE kod=i1453 LIMIT 1;
	var bulunan Urun // tek kayıt gelecek
	vt.First(&bulunan, "kod = ?", "m1070")
	fmt.Println("\nm1070 kodlu ürünün fiyatı:", bulunan.Fiyat)
	// Yukarıda m1070 kodlu bulunan ürünün fiyatını değiştir
	fmt.Println("\n*** Ürün fiyatı güncelleniyor...")
	vt.Model(&bulunan).Update("Fiyat", 2500)
	vt.First(&bulunan, "kod = ?", "m1070")
	fmt.Println("m1070 kodlu ürünün güncel fiyatı:", bulunan.Fiyat)
	// Yukarıda m1070 kodlu bulunan ürünü sil
	fmt.Println("\n*** Ürün siliniyor...")
	vt.Delete(&bulunan)
	Listele(vt)
}
