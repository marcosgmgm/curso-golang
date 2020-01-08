package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3
	fmt.Println(i)

	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2 // x = 1, y = 2
	fmt.Println(x, y)

	x, y = y, x // Trocando o valor de x e y
	fmt.Println(x, y)

}
