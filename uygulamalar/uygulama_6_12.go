// uygulama_6_12
package main

import "fmt"

func main() {
	var yaş int
	// fonskyion, yerel değişken olarak tanımlanıyor
	yaş_kontrolü := func() {
		switch {
		case yaş < 18:
			fmt.Println("Sisteme giriş için yaşınız müsait değil.")
		default:
			fmt.Println("Yaş onaylandı. ")
		}
	}
	// Kullanıcıya yaşını sor
	fmt.Print("Yaşınız: ")
	fmt.Scanln(&yaş)
	// fonksiyonu çağıralım
	yaş_kontrolü()
}
