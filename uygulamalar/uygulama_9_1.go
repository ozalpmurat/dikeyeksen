// uygulama_9_1
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	adres, _ := net.ResolveTCPAddr("tcp", "google.com:80")
	baglanti, _ := net.DialTCP("tcp", nil, adres)
	fmt.Fprintf(baglanti, "GET / HTTP/1.0\r\n\r\n")
	durum, _ := bufio.NewReader(baglanti).ReadString('\n')
	fmt.Println("Durum: ", durum)
	baglanti.Close()
}
