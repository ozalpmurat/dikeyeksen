// uygulama_9_17
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// dosya içeriğini, baytKesit değişkenine aktar:
	baytKesit, _ := ioutil.ReadFile("deneme.txt")
	// ASCII bayt kodlarını metne dönüştürüp ekrana yaz:
	fmt.Println(string(baytKesit))
}
