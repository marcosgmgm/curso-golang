package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "José", "Mateus", "Nicole"}
	imprimirAprovados(aprovados...)
}
