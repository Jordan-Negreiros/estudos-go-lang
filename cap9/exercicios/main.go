package main

import "fmt"

func main() {

	fmt.Println("----- Exercicio 01 -----")
	numeros := [5]int{1, 2, 3, 4, 5}

	for i, v := range numeros {
		fmt.Println(i, v)
	}

	fmt.Printf("%T \n", numeros)

	fmt.Println("----- Exercicio 02 -----")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T \n", slice)

	fmt.Println("----- Exercicio 03 -----")
	fmt.Println(
		slice[:3],
		slice[4:],
		slice[1:7],
		slice[2:len(slice)-1],
	)

	fmt.Println("----- Exercicio 04 -----")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)

	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	fmt.Println(x)

	fmt.Println("----- Exercicio 05 -----")

	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	w := append(z[:3], z[6:]...)

	fmt.Println(w)

	fmt.Println("----- Exercicio 06 -----")

	estados := make([]string, 26, 26)

	fmt.Println(len(estados), cap(estados))

	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estados), cap(estados))
	fmt.Println(estados)

}
