package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 //Ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("tv50: %t, tv32: %t, sorverte: %t, saudavel: %t\n", tv50, tv32, sorvete, !sorvete)
	tv50, tv32, sorvete = compras(false, true)
	fmt.Printf("tv50: %t, tv32: %t, sorverte: %t, saudavel: %t\n", tv50, tv32, sorvete, !sorvete)
	tv50, tv32, sorvete = compras(false, false)
	fmt.Printf("tv50: %t, tv32: %t, sorverte: %t, saudavel: %t\n", tv50, tv32, sorvete, !sorvete)
}
