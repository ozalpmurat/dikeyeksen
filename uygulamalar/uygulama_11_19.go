// uygulama_11_19
// https://github.com/EngoEngine/engo/tree/master/demos/text
package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	c "engo.io/engo/common"
	"image/color"
)

type YeniSahne struct{}

var (
	HizYakinlas float32 = -0.125
	HizKaydir   float32 = 700
	Genislik    int     = 500
	Yukseklik   int     = 300
)

type metin struct {
	ecs.BasicEntity
	c.RenderComponent
	c.SpaceComponent
}

func (*YeniSahne) Preload() {
	// kullanacağımız fontu yükle
	engo.Files.Load("times.ttf")
}
func (*YeniSahne) Setup(w *ecs.World) {
	w.AddSystem(&c.RenderSystem{})
	// Klavye kaydırmasını etkinleştir:
	w.AddSystem(c.NewKeyboardScroller(HizKaydir,
		engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis))
	// fare yaklaştırmasını etkinleştir:
	w.AddSystem(&c.MouseZoomer{HizYakinlas})
	// fontu tanımla
	fnt := &c.Font{
		URL:  "times.ttf",
		FG:   color.White,
		Size: 64,
	}
	fnt.CreatePreloaded()
	//ilk metni oluştur
	Metin1 := metin{}
	Metin1.RenderComponent.Drawable = c.Text{
		Font: fnt,
		Text: "Merhaba dunya",
	}
	Metin1.SetShader(c.HUDShader) // sabit dursun
	// ikinci metni oluştur
	Metin2 := metin{}
	Metin2.RenderComponent.Drawable = c.Text{
		Font: fnt,
		Text: "Bu, iki satirlik \nbir metin",
	}
	// Hepsini RenderSystem'e ekle:
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *c.RenderSystem:
			sys.Add(&Metin1.BasicEntity,
				&Metin1.RenderComponent,
				&Metin1.SpaceComponent)
			sys.Add(&Metin2.BasicEntity,
				&Metin2.RenderComponent,
				&Metin2.SpaceComponent)
		}
	}
}
func (*YeniSahne) Type() string { return "metinSahnesi" }
func main() {
	seçenekler := engo.RunOptions{
		Title:          "Metin görüntüleme",
		Width:          Genislik,
		Height:         Yukseklik,
		StandardInputs: true,
	}
	engo.Run(seçenekler, &YeniSahne{})
}
