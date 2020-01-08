package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 1234.89,
			"Guga Pereira": 12343.21,
		},
		"J": {
			"Jose Joao": 12321.21,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	for letra, funcaLetra := range funcsPorLetra {
		fmt.Printf("Funcionarios com a letra: %s\n", letra)
		for nome, salario := range funcaLetra {
			fmt.Printf("Nome: %s, Salario: %v\n", nome, salario)
		}
	}
}
