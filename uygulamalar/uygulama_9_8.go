// uygulama_9_8
package main

import (
	"fmt"
	"strings"
)

func main() {
	rot13 := func(karakter rune) rune {
		return karakter + 13
	}
	fmt.Println(strings.Map(rot13, "klmn"))
}
