package main

import "fmt"

type hotdog int // -> tipo hotdog -> tipo base inteiro, mas hotdog != int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v %T \n", x, x)

	x = int(b)
	fmt.Printf("%v %T", x, x)
}
