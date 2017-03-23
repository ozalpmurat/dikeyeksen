// uygulama_11_5
// https://github.com/gonum/plot/wiki/Example-plots
package main

import (
	Grafik "github.com/gonum/plot"
	Çizici "github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"image/color"
	"math/rand"
)

func main() {
	GrafikVerisi := RasgeleVeriler(10)
	GrafikAlani, _ := Grafik.New()
	GrafikAlani.Title.Text = "On yüz bin milyon baloncuk"
	GrafikAlani.X.Label.Text = "X"
	GrafikAlani.Y.Label.Text = "Y"
	// Grafik alanı oluşturuldu, grafiği de oluştur:
	BaloncukGrafigi, _ := Çizici.NewBubbles(GrafikVerisi,
		vg.Points(1), vg.Points(20))
	BaloncukGrafigi.Color =
		color.RGBA{R: 196, B: 128, A: 255} // renk kodları
	GrafikAlani.Add(BaloncukGrafigi)
	// Dosyaya kaydet:
	if err := GrafikAlani.Save(4*vg.Inch,
		4*vg.Inch, "baloncuk.png"); err != nil {
		panic(err)
	}
}

// x,y,z biçiminde rasgele üçlü değerler oluştur:
func RasgeleVeriler(n int) Çizici.XYZs {
	veri := make(Çizici.XYZs, n)
	for i := range veri {
		if i == 0 {
			veri[i].X = rand.Float64()
		} else {
			veri[i].X = veri[i-1].X + 2*rand.Float64()
		}
		veri[i].Y = veri[i].X + 10*rand.Float64()
		veri[i].Z = veri[i].X
	}
	return veri
}
