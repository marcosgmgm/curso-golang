package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n <= 10 && n >= 9:
		return "A"
	case n < 9 && n >= 7:
		return "B"
	case n < 7 && n >= 5:
		return "C"
	case n < 5 && n >= 3:
		return "D"
	case n < 3 && n >= 0:
		return "E"
	default:
		return "Nota inválida"

	}
}

func main()  {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.7))
	fmt.Println(notaParaConceito(7.5))
	fmt.Println(notaParaConceito(5.4))
	fmt.Println(notaParaConceito(2.3))
	fmt.Println(notaParaConceito(11.4))
}
