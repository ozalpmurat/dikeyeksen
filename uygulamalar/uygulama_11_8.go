// uygulama_11_8
// https://github.com/lazywei/go-opencv
package main

import (
	"fmt"
	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	// resim dosyasını yükle
	ResimDosyasi := "Gopher.png"
	resim := opencv.LoadImage(ResimDosyasi)
	defer resim.Release()
	// pencereyi ayarla
	pencere := opencv.NewWindow("Go-OpenCV uygulaması")
	defer pencere.Destroy()
	// ayar çubuğu yerleştir ve değerini alıp işle
	pencere.CreateTrackbar("Ayar çubuğu", 1, 100,
		func(deger int) {
			fmt.Printf("Değer = %d\n", deger)
		})
	pencere.ShowImage(resim)
	// Bir tuşa basıncaya kadar bekle
	opencv.WaitKey(0)
}
