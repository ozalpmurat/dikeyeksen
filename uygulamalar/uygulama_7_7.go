// uygulama_7_7
package main

import (
	"fmt"
)

// taşıt isminde bir yapı (tür) tanımlanıyor
type taşıt struct {
	Adı string
}

// taşıt türüne bağlı ve ilerle() isminde bir metot tanımlanıyor
func (t *taşıt) ilerle() {
	fmt.Println(t.Adı, "ileri doğru gidiyor...")
}

// deniz_taşıtı isminde yeni bir tür tanımlanıyor
type deniz_taşıtı struct {
	taşıt       // *** Gömülü Tür Tanımlandı ***
	yelken bool // deniz_taşıtı'nda, yelken isimli bir mantıksal alan olsun.
}

func main() {
	var tren taşıt
	tren.Adı = "Haydarpaşa-Adapazarı Ekspresi"
	tren.ilerle() // taşıt'ın ilerle() metotunu çağırıyoruz
	var kayık deniz_taşıtı
	kayık.Adı = "Ertuğrul" // taşıt altında tanımlanmış olan "Adı" alanı
	kayık.yelken = false   // yelkeni yok
	fmt.Println(kayık)
	fmt.Println("// --- metotlar çağırılıyor ---")
	kayık.taşıt.ilerle() // taşıt'a ait olan metotu çağırıyoruz
	kayık.ilerle()       // aynı metodu doğrudan çalıştıralım
	fmt.Println(kayık.taşıt.Adı)
	fmt.Println(kayık.Adı)
}
