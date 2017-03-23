// uygulama_7_9
package main

import "fmt"

// Özel tanımladığımız veri türlerimiz: Köpek ve Rüzgar.
// Köpek türünü yapı olarak, Rüzgar’ı ise metin olarak tanımladık.
// Özellikle farklı türler kullandık, çünkü arayüz kullanımı hepsinde aynı.
type Köpek struct {
	tür string
	yaş int
}
type Rüzgar string

// Yukarıda tanımladığımız türlere ait metotlar
func (x Köpek) SesCikarma() string {
	return "Hav Hav!"
}
func (x Rüzgar) SesCikarma() string {
	return "Fiuuvvv !"
}

// Arayüz (interface) tanım kısmı burası.
// Bu arayüz üzerinden erişilebilecek olan metotları burada belirtiriz.
// Arayüzde metodun sadece ismi yazılır. İçerik yazılmaz.
// Bu kodları başkası kullanacaksa, oluşturacağı Varlık'larda hangi
// metodları kullanabileceğini arayüz kısmından görecektir.
type Varlik interface {
	SesCikarma() string
}

// Bir arayüzü girdi olarak alan fonksiyon
func SesCikar(x Varlik) {
	fmt.Println(x.SesCikarma())
}
func main() {
	var Karabaş Köpek
	var Poyraz Rüzgar
	// SesCikar fonksiyonuna değişkenlerimizi gönderiyoruz.
	SesCikar(Karabaş)
	SesCikar(Poyraz)
	// Arayüz kullanmadan, doğrudan türe erişerek metot çağıralım:
	fmt.Println(Karabaş.SesCikarma())
}
