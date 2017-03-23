// uygulama_8_13
package main

import "time"
import "fmt"

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("2 saniye geçti")
}
