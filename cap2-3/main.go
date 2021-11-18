package main

import "fmt"

var w = "escopo de pacote" // -> só disponível dentro do pacote

func main()  {
	
	x := 10
	y := "Olá"
	z := 10.0

	var k = "escopo de função" // -> só disponível dentro da função

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)

	x = 20
	// x := 20 -> só é possível utilizar o operador := no momento da declaração da variável
	fmt.Printf("x: %v, %T\n", x, x)

	fmt.Println(w)
	fmt.Println(k)
}