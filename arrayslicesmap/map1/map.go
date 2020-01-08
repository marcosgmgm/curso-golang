package main

import "fmt"

func main() {
	//var aprovados map[int] string
	//Mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678] = "Maria"
	aprovados[11111111] = "Pedro"
	aprovados[22222223] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678]) // Imprime o value da key 12345678
	delete(aprovados, 12345678) // delete o value e ker 12345678
	fmt.Println("Apos delete")
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
