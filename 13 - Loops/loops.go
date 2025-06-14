package main

import (
	"fmt"
	"time"
)
func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i:", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j+= 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"Joao", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println( nome)
	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Kevin",
		"sobrenome": "Velez",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	fmt.Println(usuario)

	type usuarioStruct struct {
		nome string
		sobrenome string
	}
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}