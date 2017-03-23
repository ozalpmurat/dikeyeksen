package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Söz struct {
	Söyleyen string
	Zaman    int
	Cümle    string
}

func main() {
	// web sayfasını indir ve işle
	url := "http://localhost:8000"
	http_cevap, _ := http.Get(url)
	json_cevap, _ := ioutil.ReadAll(http_cevap.Body)
	defer http_cevap.Body.Close()
	// JSON elemanlarını tutacak olan yapı:
	var sözler []Söz
	// JSON'ı ayrıştır ve sözler yapı dizisine aktar
	json.Unmarshal(json_cevap, &sözler)
	// sözler dizisinin eleman sayısını yazdır
	fmt.Println("JSON'dan gelen ", len(sözler), "tane veri:\n---")
	// sözler dizisini ekrana düzenleyerek yazdır
	for i, s := range sözler {
		fmt.Printf("%d) %d sene önce, %s demiş ki: \"%s\"\n",
			i,
			time.Now().Year()-s.Zaman, // (şimdiki yıl)-(söz yılı)
			s.Söyleyen,
			s.Cümle)
	}
}
