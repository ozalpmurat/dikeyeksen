// uygulama_11_17
// https://engo.io/tutorials/01-hello-world
package main

import "engo.io/ecs"
import "engo.io/engo"

type Sahne1 struct{}

// Oyun türünü belirten metin
func (*Sahne1) Type() string {
	return "ilkOyun"
}

// Ses, resim, vb dosyaları yükle
func (*Sahne1) Preload() {}

// Setup bloğu, sahneye sistem ve
// varlık eklemek için kullanılır
func (*Sahne1) Setup(*ecs.World) {}
func main() {
	opts := engo.RunOptions{
		Title:  "Merhaba dünya",
		Width:  250,
		Height: 250,
	}
	engo.Run(opts, &Sahne1{})
}
