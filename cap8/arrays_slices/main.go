package main

import "fmt"

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	fmt.Println(len(numeros))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(len(slice))

	slice = append(slice, 7)
	fmt.Println(slice)
	fmt.Println(len(slice))

	for i, v := range slice {
		fmt.Println("Indice:", i, "Valor:", v)
	}

	slice2 := []int{55, 1, 20, 40}
	total := 0

	for _, v := range slice2 {
		total += v
	}

	fmt.Println(total)

	sliceParcial := slice2[0:2]
	totalParcial := 0

	for _, v := range sliceParcial {
		totalParcial += v
	}

	fmt.Println(totalParcial)

	sliceTeste := []int{1, 2, 3, 4, 5}
	fmt.Println(append(sliceTeste[:2], sliceTeste[3:]...)) // pega o elemento 0,1 + os elementos 3 em diante
}
