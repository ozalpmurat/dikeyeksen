// uygulama_11_16
// https://hajimehoshi.github.io/ebiten/examples/perspective.html
package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/jpeg"
)

const PencereEni = 320
const PencereBoyu = 240

var Resim *ebiten.Image

type özellikler struct {
	resim *ebiten.Image
}

func (p özellikler) Len() int {
	_, boy := p.resim.Size()
	return boy
}
func (p özellikler) Dst(i int) (x0, y0, x1, y1 int) {
	en, boy := p.resim.Size()
	genislik := en + i*3/4
	x := ((boy - i) * 3 / 4) / 2
	return x, i, x + genislik, i + 1
}
func (p özellikler) Src(i int) (x0, y0, x1, y1 int) {
	en, _ := p.resim.Size()
	return 0, i, en, i + 1
}
func Güncelle(ekran *ebiten.Image) error {
	islem := &ebiten.DrawImageOptions{
		ImageParts: &özellikler{Resim},
	}
	en, boy := Resim.Size()
	AzamiEn := float64(en) + float64(boy)*0.75
	islem.GeoM.Translate(-AzamiEn/2, -float64(boy)/2)
	islem.GeoM.Translate(PencereEni/2, PencereBoyu/2)
	ekran.DrawImage(Resim, islem)
	return nil
}
func main() {
	Resim, _, _ = ebitenutil.NewImageFromFile("gophers.jpg", 0)
	ebiten.Run(Güncelle, PencereEni,
		PencereBoyu, 2, "Perspektif (Ebiten Demo)")
}
