// uygulama_8_1

package main

import "fmt"

func f1() {

	for i := 0; i < 10; i++ {

		fmt.Print(i, " ")

	}

}

func f2() {

	harfler := "abcdefghij"

	for _, harf := range harfler {

		fmt.Print(string(harf))

	}

}

func main() {

	go f1()

	go f2()

	// programdan çıkmadan Enter’a basılmasını bekleyelim

	var EnterTuşu string

	fmt.Println("Çıkmak için Enter tuşuna basın")

	fmt.Scanln(&EnterTuşu)

}
