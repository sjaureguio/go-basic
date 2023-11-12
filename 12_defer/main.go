package main

import (
	"fmt"
	"os"
)

func main() {
	// LIFO (First Input Last Output)

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	// IMPORTANTE! Los parametros que las funciones se evaluan antes
	// de agregar a la pila

	// Caso de uso del defer
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Ocurrió un error al cerrar el archivo: %v", err)
		}
	}(file)

	_, err = file.Write([]byte("Hello Gophers \n"))
	if err != nil {
		fmt.Printf("Ocurrió un error al escribir el archivo: %v", err)
		return
	}

}
