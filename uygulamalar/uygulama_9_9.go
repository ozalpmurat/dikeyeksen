// uygulama_9_9
package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
