// uygulama_9_4
package main

import (
	"fmt"
	"strings"
)

func main() {
	istiklal_marşı := `
Korkma, sönmez bu şafaklarda yüzen al sancak;
Sönmeden yurdumun üstünde tüten en son ocak.`
	sozcukler := strings.Fields(istiklal_marşı)
	fmt.Printf("Sözcükler: %q\n", sozcukler)
	fmt.Println("Sözcük sayısı: ", len(sozcukler))
}
