// uygulama_9_11
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("a-b-c-d-e", "-"))
	fmt.Printf("%q\n", strings.SplitAfter("a-b-c-d-e", "-"))
	fmt.Printf("%q\n", strings.SplitN("a-b-c-d-e", "-", 3))
	fmt.Printf("%q\n", strings.SplitAfterN("a-b-c-d-e", "-", 3))
}
