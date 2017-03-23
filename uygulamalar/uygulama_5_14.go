// uygulama_5_14
package main

import "fmt"

func main() {
	sozluk := make(map[string]string)
	sozluk["IP"] = "Internet Protocol"
	sozluk["GNU"] = "GNU is Not Unix"
	sozluk["WWW"] = "World Wide Web"
	delete(sozluk, "GNU")
	if karşılık, sonuç := sozluk["GNU"]; sonuç {
		fmt.Println(karşılık)
	} else {
		fmt.Println("Aranan sözcük bulunamadı")
	}
}
