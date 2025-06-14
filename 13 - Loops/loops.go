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
}