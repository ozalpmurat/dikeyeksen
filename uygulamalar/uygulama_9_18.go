// uygulama_9_18
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	baytKesit, hata := ioutil.ReadFile("deneme.txt")
	if hata != nil {
		fmt.Println("dosya açılamadı")
		return
	}
	fmt.Println(string(baytKesit))
}
