package main

import "fmt"

func main() {

	// Switch clásico (NO colocar break pra romper el flujo)

	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	// Evaluar multiples expresiones
	animal := "Horse"
	switch animal {
	case "Parrot", "Eagle":
		fmt.Println("Es un ave")
	case "Horse", "Monkey":
		fmt.Println("Es un mamífero")
	}
}
