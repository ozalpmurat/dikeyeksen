// uygulama_11_15
// https://hajimehoshi.github.io/ebiten/
package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen,
		"Merhaba dunya!\nSize selam getirdim.")
	return nil
}
func main() {
	ebiten.Run(update, 160, 120, 2,
		"Pencere başlığı")
}
