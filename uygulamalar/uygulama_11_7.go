// uygulama_11_7
package main

import (
	"fmt"
	"github.com/koyachi/go-nude"
)

func main() {
	resimler := []string{
		"heykel.jpg", "lena.gif", "çift.png", "koşucu.png"}
	for _, resim := range resimler {
		ciplakmi, _ := nude.IsNude(resim)
		fmt.Println(ciplakmi, "\t", resim)
	}
}
