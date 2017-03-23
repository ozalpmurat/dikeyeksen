// uygulama_6_15
package main

import "fmt"

func main() {
	panic("Ayyy !")
	hata := recover() // bu satır çalışmayacaktır
	fmt.Println(hata)
}
