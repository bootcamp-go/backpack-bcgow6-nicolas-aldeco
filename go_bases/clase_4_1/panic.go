package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	_, err := os.ReadFile("./panic.go")
	if err != nil {
		panic(err)
	}

	// var arr []int

	// a := arr[4]

	// var a *int

	// b := 10

	// a = &b

	defer func() {
		if err := recover(); err != nil {
			log.Println("No funciono", err)
		} else {
			log.Println("Funciono correctamente")
		}
	}()
	// Situacion de panic
	// ch := make(chan int)
	// fmt.Println(<-ch)
	//

	retorno, err := FuncionImportante(0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Retorno:", retorno)

}

func FuncionImportante(estado int) (string, error) {
	if estado == 0 {
		return "", errors.New("No hay estado")
	}
	return "Si hay estado", nil
}
