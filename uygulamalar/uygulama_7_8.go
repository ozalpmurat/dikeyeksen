// uygulama_7_8
package main

import "fmt"

// taşıt isminde bir tür tanımlanıyor
type taşıt struct {
	Adı string
}

// silah isminde bir tür tanımlanıyor
type silah struct {
	Adı string
}

// taşıt türüne bağlı ve ilerle() isminde bir metot tanımlanıyor
func (t taşıt) ilerle() {
	fmt.Println(t.Adı, "ileri doğru gidiyor...")
}

// silah türüne bağlı ve ateş() isminde bir metot tanımlanıyor
func (s silah) ateş() {
	fmt.Println(s.Adı, " silahı seçildi: ATEŞ !")
}

// deniz_taşıtı isminde yeni bir tür tanımlanıyor
type deniz_taşıtı struct {
	taşıt // *** Gömülü Tür tanımladık. taşıt'ın özellikleri geldi.
	silah // *** Gömülü Tür tanımladık. savaş gemisi oluşturursak, silah gerekli.
}

func main() {
	var gemi deniz_taşıtı
	gemi.taşıt.Adı = "Nusret" // taşıt altında tanımlanmış olan "Adı" alanı
	gemi.silah.Adı = "Mayın"  // silah altında tanımlanmış olan "Adı" alanı
	fmt.Println(gemi, "gemisi")
	fmt.Println("// --- metotlar çağırılıyor ---")
	gemi.ilerle() // ilerle metodu sadece taşıt'ta var.
	gemi.ateş()   // ateş metodu sadece silah'ta var.
}
