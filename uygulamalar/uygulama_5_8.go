// uygulama_5_8
package main

import "fmt"

func main() {
	dizi := [4]string{"a", "b", "c", "d"}
	kesit := dizi[:3]
	fmt.Printf("1) eleman sayısı=%d, kapasite=%d, %v\n",
		len(kesit), cap(kesit), kesit)
	kesit = append(kesit, "x")
	fmt.Printf("2) eleman sayısı=%d, kapasite=%d, %v\n",
		len(kesit), cap(kesit), kesit)
	kesit = append(kesit, "y")
	fmt.Printf("3) eleman sayısı=%d, kapasite=%d, %v\n",
		len(kesit), cap(kesit), kesit)
	kesit = kesit[2:]
	fmt.Printf("4) eleman sayısı=%d, kapasite=%d, %v\n",
		len(kesit), cap(kesit), kesit)
}
