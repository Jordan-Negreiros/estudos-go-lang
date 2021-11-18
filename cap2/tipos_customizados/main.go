package main

import "fmt"

type hotdog int // -> tipo hotdog -> tipo base inteiro, mas hotdog != int

var b hotdog

func main() {
	fmt.Printf("%T", b)	
}