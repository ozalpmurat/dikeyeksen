// uygulama_5_10
package main

import "fmt"

func main() {
	kaynak := []int{1, 2, 9, 9}
	hedef := make([]int, 2)
	fmt.Println(kaynak, hedef)
	sayi := copy(hedef, kaynak)
	fmt.Println("kopyalanan eleman sayısı: ", sayi)
	fmt.Println(kaynak, hedef)
}
