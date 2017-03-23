// uygulama_5_6
package main

import "fmt"

func main() {
	nokta := [3][2]int{{11, 12}, {21, 22}, {31, 32}}
	fmt.Println("Dizinin tamamı: ", nokta)
	fmt.Println("\n3x2 matris biçiminde: ")
	for satir := 0; satir < 3; satir++ { // 0,1,2 indisli satırlar
		for sutun := 0; sutun < 2; sutun++ { // 0,1 indisli sütunlar
			fmt.Print(nokta[satir][sutun], " ")
		}
		fmt.Println()
	}
}
