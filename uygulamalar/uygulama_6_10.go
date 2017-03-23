// uygulama_6_10
package main

import "fmt"

func main() {
	fmt.Println(faktöriyel(4))
}
func faktöriyel(sayı int) int {
	if sayı == 0 {
		return 1
	}
	return sayı * faktöriyel(sayı-1)
}
