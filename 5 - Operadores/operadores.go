package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 4
	divisao := 10 * 5
	restodaDivisao := 10 % 3

	fmt.Println(soma, subtracao, multiplicacao, divisao, restodaDivisao)

	var n1 int16 = 10
	var n2 int16 = 25
	soma1 := n1 + n2
	fmt.Println(soma1)

	// fim dos aritmetricos

	// atribuicao
	// operdores unarios
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	numero--
	numero *= 3
	numero /= 2
	numero %= 3
	fmt.Println(numero)

	// fim dos operadores unarios

	// operadores ternarios

	var text string
	if numero > 5 {
		text = "maior que 5"
	} else {
		text = "menor que 5"
	}
	fmt.Println(text)
}
