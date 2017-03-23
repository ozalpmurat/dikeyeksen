// uygulama_9_12
package main

import (
	"fmt"
	"strings"
)

func main() {
	metin := "BeN aşığIM dalgalarıN sesinE"
	fmt.Println(strings.Title(metin))
	fmt.Println(strings.ToLower(metin))
	fmt.Println(strings.ToUpper(metin))
	fmt.Println(strings.Title(strings.ToLower(metin)))
}
