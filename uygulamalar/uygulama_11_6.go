// uygulama_11_6
package main

import (
	"fmt"
	DilTahmin "github.com/endeveit/guesslanguage"
)

func main() {
	fmt.Println(DilTahmin.Guess("Benim sadık yarim kara topraktır."))
	fmt.Println(DilTahmin.Guess("Love all, trust a few, do wrong to none."))
	fmt.Println(DilTahmin.Guess("Si quieres hacer reir a dios, cuéntale tus planes"))
	fmt.Println(DilTahmin.Guess("Steter Tropfen höhlt den Stein."))
	fmt.Println(DilTahmin.Guess("Merhaba dünya!"))
}
