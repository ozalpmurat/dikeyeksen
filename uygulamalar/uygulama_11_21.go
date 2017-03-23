// uygulama_11_21
package main

import (
	"fmt"
	ui "github.com/gizak/termui"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	ui.Init()
	defer ui.Close()
	bellek, _ := mem.VirtualMemory()
	fmt.Printf("Toplam: %dMB, Kullanılan:%dMB\n",
		bellek.Total/(1024*1024), bellek.Used/(1024*1024))
	oran := 100 * (float64(bellek.Used) / float64(bellek.Total))
	fmt.Printf("Kullanılan: %%%.0f", oran)
	mk := ui.NewPar("Çıkmak için \"q\" tuşuna basın")
	mk.BorderLabel = "Metin Kutusu"
	mk.Y = 3      // Satır numarası
	mk.Height = 3 // Genişlik
	mk.Width = 50 // Yükseklik
	bk := ui.NewGauge()
	bk.BorderLabel = "Bellek kullanımı"
	bk.Percent = int(oran)
	bk.Y = 6
	bk.Width = 50
	bk.Height = 5
	ui.Render(mk, bk) // Arayüzü oluştur
	// Çıkış için q tuşu kullanılsın:
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop() // Arayüzü göster
}
