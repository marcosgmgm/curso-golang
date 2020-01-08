package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //Operacao bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // Canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // Operacao bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
