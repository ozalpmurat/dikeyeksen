// uygulama_6_20
package main

import "fmt"

func main() {
	adres := new(int)
	*adres = 2016
	fmt.Println("Adres:", adres)
	fmt.Println("Adres'teki veri:", *adres)
}
