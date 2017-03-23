// uygulama_7_1
package main

import "fmt"

// vasıta isminde bir tür oluşturalım. Bu tür içinde 2 tamsayı, 2 metin alanı olsun.
type vasıta struct {
	yıl, fiyat int
	renk, tür  string
}

func main() {
	oto1 := vasıta{2010, 25000, "bordo", "4x4"}
	var tosbaa vasıta
	mx := vasıta{tür: "motosiklet", yıl: 2010, fiyat: 15000}
	// yukarıda üç tane yapı değişkeni oluşturuldu
	tosbaa.yıl = 1980
	tosbaa.fiyat = 10000
	tosbaa.tür = "3kapı"
	// ekrana yazdıralım
	fmt.Println(oto1)
	fmt.Println(tosbaa)
	fmt.Println(mx)
	fmt.Println("Toplam fiyat=", oto1.fiyat+tosbaa.fiyat+mx.fiyat)
}
