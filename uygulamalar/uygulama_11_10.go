// uygulama_11_10
// https://github.com/PuerkitoBio/goquery
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func TRT1_haberleri() {
	// TRT1 web sayfasına bağlan:
	belge, err := goquery.NewDocument(
		"http://www.trt1.com.tr/")
	if err != nil {
		log.Fatal(err)
	}
	// TRT1 CSS sınıfları: row & caption
	belge.Find(".row .caption").Each(func(
		i int, s *goquery.Selection) {
		baslik := s.Find("a").Text()
		metin := s.Find("p").Text()
		fmt.Printf(
			"Haber %d\n * Başlık:%s\n * Metin: %s\n",
			i+1, baslik, metin)
	})
}
func main() {
	TRT1_haberleri()
}
