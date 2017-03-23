// uygulama_5_12
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	sozluk := make(map[string]string)
	sozluk["IP"] = "Internet Protocol"
	sozluk["LAMP"] = "Linux Apache MySQL PHP"
	sozluk["GNU"] = "GNU is Not Unix"
	sozluk["WWW"] = "World Wide Web"
	sozluk["IMAP"] = "Internet Message Access Protocol"
	tum_elemanlar, _ := json.MarshalIndent(sozluk, "", " ")
	fmt.Println(string(tum_elemanlar))
}
