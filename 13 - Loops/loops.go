package main

import (
	"fmt"
	// "time"
)
func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i:", i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j+= 5 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println( nome)
	}
}