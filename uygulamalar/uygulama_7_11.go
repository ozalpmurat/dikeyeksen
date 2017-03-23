// uygulama_7_11
package main

import "fmt"

type Köpek struct {
	tür string
	ses string
}

func (k Köpek) SesCikarma() {
	fmt.Println(k.ses)
}

type Varlik interface {
	SesCikarma()
}

func SesCikar(x Varlik) {
	x.SesCikarma()
}
func main() {
	var Karabaş Köpek
	Karabaş = Köpek{tür: "Kangal", ses: "Hav hav!"}
	SesCikar(Karabaş)
	Karabaş.SesCikarma()
	// fmt.Println(Karabaş.SesCikarma())
}
