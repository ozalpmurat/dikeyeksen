// uygulama_7_3
package main

import "fmt"

type kuşlar struct {
	tür, renk, isim string
	yaş             int
}

func main() {
	var papağan [3]kuşlar
	papağan[0] = kuşlar{yaş: 1, renk: "beyaz", isim: "Çaki", tür: "Sultan"}
	papağan[1].yaş = 3
	papağan[1].renk = "mavi"
	papağan[1].tür = "Yeşil_papağan"
	papağan[1].isim = "İnci"
	// ekrana yazdıralım
	for _, i := range papağan {
		fmt.Println(i)
	}
}
