package main

import (
	"fmt"
)

var horas int = 10

const segundos = 1

func main() {
	minutos := 25.9
	fmt.Printf("segundos:\tValor: %d, Tipo de dato: %T\n", segundos, segundos)
	fmt.Printf("minutos:\tValor: %f, Tipo de dato: %T\n", minutos, minutos)
	fmt.Printf("horas:\tValor: %d, Tipo de dato: %T\n", horas, horas)
}
