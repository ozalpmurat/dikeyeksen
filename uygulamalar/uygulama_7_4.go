// uygulama_7_4
package main

import "fmt"

type insan struct {
	boy float32
}

func (birisi insan) idealkilo() float32 {
	kilo := 0.9*birisi.boy - 85
	return kilo
}
func main() {
	murat := insan{boy: 172}
	fmt.Println(murat.idealkilo())
}
