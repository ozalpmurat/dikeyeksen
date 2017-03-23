// uygulama_7_6
package main

import "fmt"

type SAYI int // SAYI, tamsayı türündedir
func (x SAYI) karesi() int {
	return int(x * x)
}
func main() {
	var a, b SAYI = 5, 2
	c := a - b
	// a, b ve c SAYI türünde değişkenlerdir
	fmt.Println(a.karesi(), c.karesi())
}
