package main

import "fmt"

func main() {
	fmt.Println(obterResultado(6.2))
}

// Nao tem o operador ternario em Go
func obterResultado(nota float64) string {
	//return nota >= 6 ? "Aprovado" : "Reprovado"
	if (nota >= 6) {
		return "Aprovado"
	}
	return "Reprovado"
}
