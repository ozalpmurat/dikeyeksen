// uygulama_5_11
package main

import "fmt"

func main() {
	il_nufus := make(map[string]int)
	il_nufus["Bilecik"] = 210000
	il_nufus["İstanbul"] = 15000000
	il_nufus["Sakarya"] = 1000000
	fmt.Println(il_nufus["Bilecik"])
	fmt.Println(il_nufus)
}
