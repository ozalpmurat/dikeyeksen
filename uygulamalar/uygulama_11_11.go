// uygulama_11_11
// https://github.com/stianeikeland/go-rpio
package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

var (
	led_pin = rpio.Pin(2) // GPIO-2 pini
)

func main() {
	// GPIO portunu aç
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// GPIO portunu kapatmayı unutma:
	defer rpio.Close()
	// pin çıkış için ayarlansın
	led_pin.Output()
	// pin’in durumunu (açık/kapalı) 10 kere değiştir
	for i := 0; i < 10; i++ {
		led_pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
