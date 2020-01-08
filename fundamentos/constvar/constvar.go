package main

import (
	"fmt"
	m "math" //Renomeando o math
)

func main()  {
	const PI float64 = 3.1415
	var raio = 3.2 // Tipo (Float64) Inferido pelo compilador

	//Forma reduzida de criar um var
	area := PI * m.Pow(raio, 2)

	fmt.Printf("A área da circunferência é: %v \n", area)

	//Outra forma de declarar constante

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//Atribuicao em uma unica linha

	var e, f bool = true, false

	fmt.Println("e:", e, "f:", f)

	//Atribuicao em uma unica linha com declaracao reduzida
	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
