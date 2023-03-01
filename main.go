package main

import (
	"adult_go/validator"
	"fmt"
)

func main() {
	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Println("Por favor digite su edad")
	fmt.Scanln(&eleccion)

	result := validator.Is_welcome(uint(eleccion))
	fmt.Println(result)
}
