// uygulama_8_15
package main

import "time"
import "fmt"

func main() {
	// 1 saniyelik zaman çizelgesi oluşturalım
	peryot := time.NewTicker(time.Second)
	go func() {
		fmt.Println("Saniyede 1 tane iş yapabilirim:")
		// Her bir peryot için çalışan döngü
		for zaman := range peryot.C {
			fmt.Println(zaman)
		}
	}()
	time.Sleep(time.Second * 4)
	peryot.Stop()
}
