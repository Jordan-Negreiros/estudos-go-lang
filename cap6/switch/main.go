package main

import "fmt"

func main() {

	estaonotrabalhohoje := "maria"

	switch estaonotrabalhohoje {

	case "maria", "ze":
		fmt.Println("Hoje é dia do Time 1")
	case "aria", "lia":
		fmt.Println("Hoje é dia do Time 2")
		fallthrough
	case "exemplo de fallthrough":
		fmt.Println("fallthrough -> é executando se o case anterior for verdadeiro")
	default:
		fmt.Println("folga")

	}

	x := 10
	y := 10

	fmt.Println(x == 10 || x == 11)
	fmt.Println(x == 10 && y == 11)
	fmt.Println(x != 10)
	fmt.Println(!true || x == 11)
}
