package main

import (
	"fmt"
)

func main() {
	// contoh 1
	var message = `Nama saya "Zulfadli Rizal". Salam kenal . Mari belajar "Golang".`

	fmt.Println(message)

	// contoh 2
	var name = "Zulfadli Rizal"
	var age = "26"
	var sentence = `halo nama saya "` + name + `" dan berumur "` + age + `"`

	fmt.Println()
	fmt.Println(sentence)
}
