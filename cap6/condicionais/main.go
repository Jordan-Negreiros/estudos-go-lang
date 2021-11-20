package main

import "fmt"

func main() {

	if x := 10; x < 10 {
		fmt.Printf("x < 10")
	} else if x > 10 {
		fmt.Printf("x > 10")
	} else {
		fmt.Printf("x == 10")
	}
}
