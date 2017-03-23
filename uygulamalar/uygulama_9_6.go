// uygulama_9_6
package main

import (
	"fmt"
	"strings"
)

func main() {
	metin := "Korkma, sönmez bu şafaklarda yüzen al sancak;"
	fmt.Println(strings.Index(metin, "k"))
	fmt.Println(strings.IndexAny(metin, "?.,;!:"))
}
