// uygulama_6_8
package main

import "fmt"

func main() {
	isim := isim_al()
	yaş := yaş_al()
	fmt.Printf("\nHoşgeldiniz %s, yaşınız: %d. ", isim, yaş)
	kontrol_et(yaş)
}
func isim_al() (isim string) {
	fmt.Print("İsminiz: ")
	fmt.Scanln(&isim)
	return
}
func yaş_al() (yaş int) {
	fmt.Print("Yaşınız: ")
	fmt.Scanln(&yaş)
	return
}
func kontrol_et(yaş int) {
	switch {
	case yaş < 18:
		fmt.Println("Sisteme giriş için yaşınız müsait değil. ")
	case yaş < 120:
		fmt.Println("Sisteme başarılı olarak giriş yaptınız. ")
		oturum_aç()
	default:
		fmt.Println("Şaka yapıyorsunuz değil mi? ")
	}
}
func oturum_aç() { fmt.Println("İşiniz bitince oturumunuzu kapatmayı unutmayın.") }
