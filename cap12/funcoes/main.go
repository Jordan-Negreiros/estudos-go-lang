package main

import "fmt"

type Empregado interface {
	imprimirNome(nome string)
	imprimirSalario(base int, taxa int) int
}

type Emp int

func (e Emp) imprimirNome(nome string) {
	fmt.Println("Empregado Id: \t", e)
	fmt.Println("Empregado Nome: \t", nome)
}

func (e Emp) imprimirSalario(base int, taxa int) int {
	var salario = (base * taxa) / 100
	return base - salario
}

func main() {
	var e1 Empregado
	e1 = Emp(1)
	e1.imprimirNome("Nome")
	fmt.Println("Salário:", e1.imprimirSalario(20000, 5))
}
