// uygulama_7_10
package main

import "fmt"

// Tür tanımları
type Köpek struct {
	tür string
	yaş int
}
type Telefon string
type Rüzgar string

// Türlere ait metotlar
func (x Köpek) SesCikarma() string {
	return "Hav Hav!"
}
func (x Telefon) SesCikarma() string {
	return "Didit didit!"
}
func (x Rüzgar) SesCikarma() string {
	return "Fiuuvvv !"
}
func (x Rüzgar) YağmurGetir() string {
	return "Bazı insanlar yağmuru hisseder, diğerleri sadece ıslanır."
}

// Arayüz tanımları
type Varlik interface {
	SesCikarma() string
}
type Meteorolojik interface {
	YağmurGetir() string
}

// Arayüz-metot aracılığı yapan fonksiyonlar
func SesCikar(x Varlik) {
	fmt.Println(x.SesCikarma())
}
func YağmurGetir(x Meteorolojik) {
	fmt.Println(x.YağmurGetir())
}
func main() {
	// Değişkenler (Nesneler gibi de düşünebiliriz)
	var Karabaş Köpek
	Karabaş.tür = "Kangal"
	Karabaş.yaş = 3
	var CepTelefonu Telefon
	var Poyraz Rüzgar
	// Arayüz üzerinden metotlara erişim
	SesCikar(Karabaş)
	SesCikar(CepTelefonu)
	SesCikar(Poyraz)
	YağmurGetir(Poyraz)
	// Arayüz kullanmadan metot çağıralım
	fmt.Println("--------")
	fmt.Println(Karabaş.SesCikarma())
	fmt.Println(Poyraz.YağmurGetir())
}
