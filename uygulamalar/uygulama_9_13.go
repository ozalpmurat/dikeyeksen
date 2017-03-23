// uygulama_9_13
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	tr := unicode.TurkishCase
	metin := "BEN aşığım dalgaların SESİNE"
	fmt.Println(strings.ToLowerSpecial(tr, metin))
	fmt.Println(strings.ToUpperSpecial(tr, metin))
}
