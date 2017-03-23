// uygulama_11_4
package main

import (
	"fmt"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// main bloğu içinde grafikleri oluşturuyoruz
func main() {
	const (
		GrafikEni  = 800 // bir grafik için
		GrafikBoyu = 500 // bir grafik için
		Sutun      = 2
		Satir      = 1
	)
	// *** Pasta grafiği ***
	PastaGrafigi := chart.PieChart{Title: "Birşeylerin dağılımı"}
	PastaGrafigi.AddDataPair("Veriler",
		[]string{"Aaa", "Bbb", "Ccc", "Ççç", "Ddd", "Eee"},
		[]float64{10, 20, 30, 35, 15, 25})
	PastaGrafigi.Data[0].Samples[3].Flag = true
	PastaGrafigi.Inner = 0.5
	PastaGrafigi.Key.Border = -1
	PastaGrafigi.FmtVal = chart.AbsoluteValue
	// *** Sinüs eğrisi ***
	SinusGrafigi := chart.ScatterChart{Title: "Sinüs grafiği"}
	SinusGrafigi.XRange.Fixed(0, 4*math.Pi, math.Pi)
	SinusGrafigi.YRange.Fixed(-1.25, 1.25, 0.5)
	SinusGrafigi.XRange.TicSetting.Format = func(f float64) string {
		w := int(180*f/math.Pi + 0.5)
		return fmt.Sprintf("%d°", w)
	}
	SinusGrafigi.AddFunc("Sin(x)",
		func(x float64) float64 { return math.Sin(x) }, chart.PlotStyleLines,
		chart.Style{Symbol: '@', LineWidth: 2,
			LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}, LineStyle: 0})
	SinusGrafigi.XRange.TicSetting.Tics,
		SinusGrafigi.YRange.TicSetting.Tics = 1, 1
	SinusGrafigi.XRange.TicSetting.Mirror,
		SinusGrafigi.YRange.TicSetting.Mirror = 2, 2
	SinusGrafigi.XRange.TicSetting.Grid,
		SinusGrafigi.YRange.TicSetting.Grid = 2, 1
	SinusGrafigi.XRange.ShowZero = true
	// Grafik dosyalarını kaydet
	pasta := GrafikKayitcisi("PastaGrafikDosyasi", 1, 1, GrafikEni, GrafikBoyu)
	pasta.GrafikCiz(&PastaGrafigi)
	sinus := GrafikKayitcisi("SinüsGrafikDosyasi", 1, 1, GrafikEni, GrafikBoyu)
	sinus.GrafikCiz(&SinusGrafigi)
}

type GrafikYapisi struct {
	N, M, W, H, Cnt int
	I               *image.RGBA
	imgFile         *os.File
}

func GrafikKayitcisi(name string, n, m, w, h int) *GrafikYapisi {
	dumper := GrafikYapisi{N: n, M: m, W: w, H: h}
	dumper.imgFile, _ = os.Create(name + ".png")
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(),
		bg, image.ZP, draw.Src)
	return &dumper
}
