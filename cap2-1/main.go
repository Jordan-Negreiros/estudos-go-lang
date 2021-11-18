package main

import "fmt"

func main() {
	// numerodebytes, erros := fmt.Println("Hello World", "Hello", 100)
	_, erros := fmt.Println("Hello World", "Hello", 100) // _ -> ignora a variável
	fmt.Println(erros)

	x := 16
	y := "string"
	z := true

	fmt.Println(x, y, z)
}