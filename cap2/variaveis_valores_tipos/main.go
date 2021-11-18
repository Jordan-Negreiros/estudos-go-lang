package main

import "fmt"

var w = "escopo de pacote" // -> só disponível dentro do pacote

var inteiro int // -> valor inicial 0 
var strings string // -> valor inicial ""
var booleano bool // -> valor inicial false
var float float64 // -> valor inicial 0.0

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

	fmt.Printf("%v, %T\n", inteiro, inteiro)
	fmt.Printf("%v, %T\n", strings, strings)
	fmt.Printf("%v, %T\n", booleano, booleano)
	fmt.Printf("%v, %T\n", float, float)
}