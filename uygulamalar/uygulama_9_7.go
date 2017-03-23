// uygulama_9_7
package main

import (
	"fmt"
	"strings"
)

func main() {
	sözcükler := []string{"süt", "bal", "yumurta"}
	fmt.Println(strings.Join(sözcükler, ", "))
	fmt.Println(strings.Join(sözcükler, "\n"))
}
