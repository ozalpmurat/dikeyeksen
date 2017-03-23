// uygulama_11_22
package main

import (
	"fmt"
	ui "github.com/gizak/termui"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

const Guncelleme = 2 //Yenileme sıklığı
func bellek() int {
	bellek, _ := mem.VirtualMemory()
	oran := 100 * (float64(bellek.Used) / float64(bellek.Total))
	return int(oran)
}

// Bilgisayarın açık kalma süresini verir
func uptime() string {
	bilgi, _ := host.Info()
	//string’e çevirip döndür:
	return fmt.Sprint(bilgi.Uptime)
}
func main() {
	ui.Init()
	defer ui.Close()
	bk := ui.NewGauge()
	bk.BorderLabel = "Bellek kullanımı"
	bk.Percent = bellek()
	bk.Y = 3
	bk.Width = 50
	bk.Height = 5
	cs := ui.NewPar(uptime())
	cs.Y = 0
	cs.Height = 3
	cs.Width = 50
	cs.BorderLabel = "Bilgisayarın çalışma süresi (uptime)"
	ui.Render(bk, cs)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	// Zamanlayıcı (timer) kısmı:
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		if t.Count%Guncelleme == 0 {
			// Guncelleme peryodunda şunları yap:
			cs.Text = uptime()
			bk.Percent = bellek()
			ui.Render(bk, cs)
		}
	})
	ui.Loop() // Arayüzü göster
}
