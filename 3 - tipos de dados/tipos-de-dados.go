package main
import (
	"fmt"
	"errors"
)
func main() {
	numero := 100000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	var numero64 byte = 123
	fmt.Println(numero64)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1200003.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234333.87
	fmt.Println(numeroReal3)

	// fim numeros reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := '%'
	fmt.Println(char)

	// fim strings

	var texto int16 
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}