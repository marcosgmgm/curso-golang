package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i //Pegando o endereço da variavel
	*p++ //Desreferenciado (pegando o valor)
	i++

	//Go nao tem aritimetica de ponteiro
	//p++

	fmt.Println(i, *p, p, &i)
}
