package main

import "fmt"

func main() {
	// Declaración y asignación de variable clasico
	var banana string
	banana = "Plátano"

	// Declaración y asignación en una linea (NO es necesario el tipo)
	var apple = "Manzana"

	// Declaración de variables con agrupamiento
	var (
		lemon string
		melon string
	)
	lemon = "Limón"
	melon = "Melón"

	// Declaración de variables corta (:)
	monkey, horse, eagle := "Mono", "Caballo", "Aguila"

	fmt.Println("Hello, World! ", banana, apple, lemon, melon)
	fmt.Println(monkey, horse, eagle)

	// Variable que acepta cualquier tipo de valor
	var x any
	fmt.Println(x)

}
