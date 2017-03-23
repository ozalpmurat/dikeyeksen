// uygulama_9_3
package main

import (
	"fmt"
	"strings"
)

func main() {
	istiklal_marşı := `
Korkma, sönmez bu şafaklarda yüzen al sancak;
Sönmeden yurdumun üstünde tüten en son ocak.
O benim milletimin yıldızıdır, parlayacak;
O benimdir, o benim milletimindir ancak. `
	boşluk_sayısı := strings.Count(istiklal_marşı, " ")
	satır_sayısı := strings.Count(istiklal_marşı, "\n")
	fmt.Println("Metindeki sözcük sayısı =", boşluk_sayısı+satır_sayısı-1)
}
