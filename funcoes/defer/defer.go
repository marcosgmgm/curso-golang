package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim metodo")
	if numero > 5000 {
		fmt.Println("Valor Alto ....")
		return 5000
	}
	fmt.Println("Valor Baixo...")
	return numero
}

func main() {
	defer fmt.Println("Fim da execucao principal")
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
