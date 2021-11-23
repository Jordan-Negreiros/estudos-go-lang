package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome      string
	idade     int
	profissao profissao
}

type profissao struct {
	titulo  string
	salario float32
}

func main() {

	cliente1 := cliente{"Ju", "Silva", true}
	cliente2 := cliente{"Lu", "Ferreira", false}
	cliente3 := cliente{"Nu", "Santos", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
	fmt.Println(cliente3)

	pessoa := pessoa{
		nome:  "Jão",
		idade: 20,
		profissao: profissao{
			titulo:  "Atendente",
			salario: 2000.10,
		},
	}

	fmt.Println(pessoa)
}
