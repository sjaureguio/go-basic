package main

import "fmt"

type User struct {
	ID       uint64
	Name     string
	Password string
}

func main() {
	var a uint8 = 220
	var ap = &a

	// Operador de dirección => &
	// Operador de desreferenciación => *

	fmt.Printf("Address: %p, Value: %v\n", ap, *ap)

	// Modifiva el valor al que apunta el puntero
	*ap = 250
	fmt.Printf("Address: %p, Value: %v\n", ap, *ap)

	user := User{
		ID:       1,
		Name:     "Solano Jauregui",
		Password: "123abc",
	}
	fmt.Printf("Address: %p, Value: %+v\n", &user, user)

	// Creando una variable tipo puntero para la estructura
	userp := &user

	// Modificar un campo de la estruc. mediante el puntero
	(*userp).Name = "Satia" // Forma compleja
	userp.Name = "Harvey"   // Forma simple

	// Función con parámetro por referencia
	sum(ap)

	fmt.Printf("Valor actualizado: %v \n", a)
}

func sum(num *uint8) {
	*num += 2
}
