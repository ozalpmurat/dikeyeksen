// uygulama_3_6
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 3
	var b = 5
	var c = "Go"
	fmt.Println("Toplam:", a+b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(c))
}
