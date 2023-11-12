package main

import "fmt"

// CONSTANTES
// No es necesario usar en su totalidad.
// Por lo tanto, no genera errores al compilar

// Constantes globales (scope package)
const (
	work  = "Dev"
	skill = "Backend"
)

// Constant iota
const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
)

func main() {
	// Declaración de constantes clásico
	const business string = "Sislii"

	// Declaración de constante mutliple
	const os, domain string = "Linux", "sislii.pe"

	// Declaración de constantes con agrupamiento
	const (
		work  = "Dev"
		skiil = "Backend"
	)

	fmt.Println(business, os, domain)
	fmt.Println(Jun)
}
