package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {

		if horas < 10 {
			fmt.Printf("horas: 0%v \n", horas)
		} else {
			fmt.Printf("horas: %v \n", horas)
		}

		for minutos := 0; minutos <= 60; minutos++ {

			if minutos < 10 {
				fmt.Printf("%v:0%v  ", horas, minutos)
			} else {
				fmt.Printf("%v:%v  ", horas, minutos)
			}

		}

		fmt.Printf("\n")
	}
}
