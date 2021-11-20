package main

import (
	"fmt"
	"time"
)

func main() {

	// exercicio 01
	for i := 1; i <= 1000; i++ {
		fmt.Printf("%v ", i)
	}

	// exercicio 02
	for i := 65; i <= 90; i++ {
		fmt.Println(i)

		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	// exercicio 03
	idade := -1
	for i := 1995; i <= time.Now().Year(); i++ {
		fmt.Println(i)
		idade += 1
	}

	fmt.Println(idade)

	// exercicio 05
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}
