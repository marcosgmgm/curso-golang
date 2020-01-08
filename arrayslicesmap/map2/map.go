package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64 {
		"José Joao": 1234.54,
		"Gabriela Silva": 12313.32,
		"Pedro Junior": 1200.0,
	}
	funcsESalarios["Camionete"] = 12000.00
	delete(funcsESalarios, "Ze ninguem") // Nao faz nada quando manda deletar algo que nao existe

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
