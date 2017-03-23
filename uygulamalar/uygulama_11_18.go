// uygulama_11_18
// https://github.com/EngoEngine/engo/tree/master/demos/edgescroller
package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"engo.io/engo/demos/demoutils"
	"image/color"
)

type Sahne1 struct{}

func (*Sahne1) Preload() {}

var (
	KenarBoslugu float32 = 100
	KaydirmaHizi float32 = 150
	Genislik     int     = 300
	Yukseklik    int     = 300
)

// Oyun başlamadan önce çalışacak kısım:
func (*Sahne1) Setup(w *ecs.World) {
	w.AddSystem(&common.RenderSystem{})
	// Kaydırmayı sağlayan kısım:
	w.AddSystem(&common.EdgeScroller{
		KaydirmaHizi, KenarBoslugu})
	// Zemini oluştur:
	demoutils.NewBackground(w, Genislik, Yukseklik,
		color.RGBA{255, 255, 255, 0},
		color.RGBA{255, 255, 255, 128})
}
func (*Sahne1) Type() string { return "2b-kaydırma" }
func main() {
	opts := engo.RunOptions{
		Title:  "Sahne kaydırma örneği",
		Width:  Genislik,
		Height: Yukseklik,
	}
	engo.Run(opts, &Sahne1{})
}
