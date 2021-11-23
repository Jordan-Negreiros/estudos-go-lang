package main

import "fmt"

func main() {
	contatos := map[string]int{
		"jão": 22223333,
		"zoe": 12345432,
	}

	fmt.Println(contatos)
	fmt.Println(contatos["zoe"])
	fmt.Println(contatos["jão"])

	if valor, existe := contatos["maria"]; !existe {
		fmt.Println(valor, existe, "Não existe")
	} else {
		fmt.Println(valor, existe, "Existe")
	}

	exemploMap := map[int]string{
		1: "um",
		2: "dois",
		3: "tres",
	}

	for key, value := range exemploMap {
		fmt.Println(key, value)
	}

	delete(exemploMap, 2)

	for key, value := range exemploMap {
		fmt.Println(key, value)
	}
}
