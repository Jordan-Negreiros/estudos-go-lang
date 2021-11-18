package main

import "fmt"

// Exercicio 02
var a int
var b string
var c bool

// Exercicio 03
var d int = 42
var f bool = true
var e string = "James Bond"

// Exercico 04
type tipo int

var g tipo

// Exercicio 05
var h int

func main() {
	// Exercicio 01
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x, y, z)

	// Exercicio 02
	fmt.Println(a, b, c) // valores iniciais (valores "zero")

	// Exercicio 03
	s := fmt.Sprintf("%v \t %v \t %v", d, f, e)
	fmt.Println(s)

	// Exercico 04
	fmt.Printf("valor -> %v, tipo -> %T \n", g, g)
	g = 42
	fmt.Printf("valor -> %v, tipo -> %T \n", g, g)

	// Exercicio 05
	h = int(g)
	fmt.Printf("valor -> %v, tipo -> %T \n", h, h)
}
