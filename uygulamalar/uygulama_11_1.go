// uygulama_11_1
package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"os"
)

func main() {
	ui.RunEx(os.Args, AnaPencere)
}
func AnaPencere() {
	// İlk düğmenin ismi, Dugme1
	Dugme1 := ui.NewPushButton()
	Dugme1.SetText("Düğme 1")
	// İkinci düğmenin ismi, DugmeTemizle
	DugmeTemizle := ui.NewPushButton()
	DugmeTemizle.SetText("Temizle")
	// Metin kutusunu oluştur:
	MetinKutusu := ui.NewPlainTextEdit()
	MetinKutusu.SetReadOnly(true)
	MetinKutusu.AppendPlainText(fmt.Sprintf("Merhaba Dünya!\nÜstteki düğmelerideneyebilirsiniz."))
	// Dugme1 basıldıysa:
	Dugme1.OnClicked(func() {
		for i := 0; i < 5; i++ {
			MetinKutusu.AppendPlainText(fmt.Sprintf("Sıra: %d", i))
		}
	})
	// DugmeTemizle basıldıysa:
	DugmeTemizle.OnClicked(func() {
		MetinKutusu.Clear()
	})
	// Yatay düğme satırını oluştur:
	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(Dugme1)
	hbox.AddWidget(DugmeTemizle)
	// Düşey satırları oluştur:
	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddWidget(MetinKutusu)
	// Elemanları ekranda göster
	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.Show()
}
