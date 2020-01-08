package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Pode adicionar mais métodos
}

type bmw7 struct {}

func (b bmw7) ligarTurbo()  {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var a esportivoLuxuoso = bmw7{}
	a.ligarTurbo()
	a.fazerBaliza()
}