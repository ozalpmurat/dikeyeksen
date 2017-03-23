// uygulama_9_10
package main

import (
	"fmt"
	"strings"
)

func main() {
	metin := "kedileri çok severim, sokaktaki kedilere yiyecek veririm"
	fmt.Println(strings.Replace(metin, "kedi", "kedi ve köpek", 1))
	fmt.Println(strings.Replace(metin, "kedi", "köpek", -1))
}
