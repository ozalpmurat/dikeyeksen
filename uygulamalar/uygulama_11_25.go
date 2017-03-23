// uygulama_11_25
package main

import (
	"fmt"
	g "github.com/soniah/gosnmp"
	"log"
)

func main() {
	g.Default.Target = "demo.snmplabs.com"
	g.Default.Community = "public"
	hata := g.Default.Connect()
	if hata != nil {
		log.Fatalf("Bağlantı hatası: %v", hata)
	}
	defer g.Default.Conn.Close()
	// Verileri sunucudan alma kısmı:
	oids := []string{"1.3.6.1.2.1.1.4.0",
		"1.3.6.1.2.1.1.6.0", "1.3.6.1.2.1.1.1.0"}
	result, hata := g.Default.Get(oids) // Değerleri al
	if hata != nil {
		log.Fatalf("Get() komutunda hata: %v", hata)
	}
	// Verileri ekrana yazdır:
	for _, variable := range result.Variables {
		fmt.Printf("%s \t %s \n",
			variable.Name, variable.Value)
	}
}
