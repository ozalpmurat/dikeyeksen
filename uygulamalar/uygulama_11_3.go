// uygulama_11_3
package main

import (
	"fmt"
	"github.com/goml/gobrain"
	"math/rand"
)

func main() {
	// Rasgele sayı üretecini besle
	rand.Seed(0)
	// İki giriş, bir çıkışlı XOR sistemi
	Desenler := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}
	// Yapay sinir ağı oluşturulsun.
	// 2 giriş, 2 gizli katman, 1 çıkışı olsun.
	YSA := &gobrain.FeedForward{}
	YSA.Init(2, 2, 1)
	fmt.Printf("%+v\n", *YSA)
	// XOR desenlerini kullanarak ağı eğit. 4000 epoch çalışsın.
	// Öğrenme katsayısı 0,6 ve momentum katsayısı 0,4 olsun.
	YSA.Train(Desenler, 4000, 0.6, 0.4, true)
	fmt.Printf("%+v\n", *YSA)
	fmt.Println("\nSistem cevapları:")
	YSA.Test(Desenler)
}
