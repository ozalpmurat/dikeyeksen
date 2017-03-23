// uygulama_9_24
// matematik.go
package matematik

// özyinelemeli faktöryel fonksiyonu
func Faktöryel(n uint64) (sonuç uint64) {
	if n > 0 {
		sonuç = n * Faktöryel(n-1)
		return sonuç
	}
	return 1
}
