// uygulama_11_14
// $GOPATH/src/azul3d.org/examples/azul3d_clearing
package main

import (
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/window"
	"image"
)

func gfxLoop(w window.Window, d gfx.Device) {
	for {
		// tüm alanı temizle (beyaz)
		d.Clear(d.Bounds(), gfx.Color{1, 1, 1, 1})
		// Dörtgenleri çiz
		// gfx.Color{R,G,B,A} Kırmızı, Yeşil, Mavi
		d.Clear(image.Rect(50, 50, 750, 400),
			gfx.Color{0, 0.5, 0.5, 1})
		d.Clear(image.Rect(150, 20, 500, 300),
			gfx.Color{1, .5, .5, 1})
		d.Clear(image.Rect(350, 150, 600, 200),
			gfx.Color{1, 1, 0, 1})
		// Görüntüyü oluştur
		d.Render()
	}
}
func main() {
	window.Run(gfxLoop, nil)
}
