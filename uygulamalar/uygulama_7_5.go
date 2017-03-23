// uygulama_7_5
package main

import "fmt"

// (x1,y1) +------+
// | |
// +------+ (x2,y2)
type Dörtgen struct {
	x1, y1, x2, y2 float64
}

func (D *Dörtgen) alan() float64 {
	kenar1 := D.x2 - D.x1
	kenar2 := D.y2 - D.y1
	return kenar1 * kenar2
}
func main() {
	Kare := Dörtgen{5, 5, 10, 10}
	fmt.Println("Kare", Kare.alan(), "birimkare")
	Dikdörtgen := Dörtgen{0, 0, 50, 10}
	fmt.Println("Dikdörtgen", Dikdörtgen.alan(), "birimkare")
}
