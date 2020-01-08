package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //Busca algum case que bate igual a switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
