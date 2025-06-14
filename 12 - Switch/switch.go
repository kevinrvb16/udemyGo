package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
		case numero ==1
	}
}
func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(50)
	fmt.Println(dia)

}
