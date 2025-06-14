package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel2++
	fmt.Println(variavel1, variavel2)

	// ponteiro eh uma referencia de memoria
	var variavel3 int
	var ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) // nil

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
