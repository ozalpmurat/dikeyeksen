// uygulama_5_9
package main

import "fmt"

func main() {
	ilk := []int{1, 4}
	son := []int{5, 3}
	butun := append(ilk, son...)
	fmt.Printf("%v\n", butun)
}
