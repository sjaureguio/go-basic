package main

import (
	"fmt"
	"slices"
)

func main() {
	// Los Slices en GO no poseen datos, son apuntadores
	// a un array

	// En Go, un slice es una estructura de datos que
	// proporciona una vista dinámica de un array.
	// A diferencia de un array, un slice no tiene
	// un tamaño fijo y puede crecer o reducirse durante
	// la ejecución del programa

	animals := [4]string{"Horse", "Donkey", "Monkey", "Parrot"}

	// Creación del un slice desde un array
	domestic := animals[:2]
	salvaje := animals[2:]
	fmt.Println("Domestics:", domestic)
	fmt.Println("Salvajes:", salvaje)

	// Slice literal
	fruits := []string{"Pinea", "Pear"}
	fmt.Println(fruits)

	// Otra forma de construir slices con make() indicando el
	// tamaño y la capacidad
	accesories := make([]string, 0, 3)

	// Agregando elementos al slice
	accesories = append(accesories, "Mouse", "Monitor", "Keyboard")

	fmt.Printf("Tamaño: %v \n", len(accesories))
	fmt.Printf("Capacidad: %v \n", cap(accesories))

	// Valor cero de un slice (nil)
	var components []string
	fmt.Println(components == nil) // true

	// copy de un slice (Destino, origen)
	c := make([]string, len(accesories))
	copy(c, accesories)

	fmt.Println(c)

	// Comparar si dos slices son iguales
	if slices.Equal(accesories, c) {
		fmt.Println("components == c")
	}

}
