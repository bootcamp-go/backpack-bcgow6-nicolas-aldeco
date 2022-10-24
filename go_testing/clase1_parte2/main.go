package main

import (
	"fmt"

	"github.com/backpack-bcgow6-nicolas-aldeco/go_testing/pkg/operaciones"
)

func main() {
	a, b := 10, 5
	sum, _ := operaciones.Add(a, b)
	fmt.Println(sum)
}
