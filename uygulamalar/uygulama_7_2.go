// uygulama_7_2
package main

import "fmt"

type insan struct {
	boy, kilo float32
	cinsiyet  string
}

func main() {
	nigar := insan{boy: 165, cinsiyet: "kadın"}
	murat := insan{boy: 172, cinsiyet: "erkek"}
	ideal_kilo(&murat)
	ideal_kilo(&nigar)
	fmt.Println(murat.kilo)
	fmt.Println(nigar.kilo)
}
func ideal_kilo(birisi *insan) {
	switch birisi.cinsiyet {
	case "erkek":
		birisi.kilo = 0.9*birisi.boy - 85
	case "kadın":
		birisi.kilo = 0.9*birisi.boy - 90
	default:
		fmt.Println("Cinsiyet alanında sorun var")
	}
}
