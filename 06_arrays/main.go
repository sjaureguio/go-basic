package main

import "fmt"

func main() {
	// UN array es una estructura de datos que almacen
	// valores del mismo tipo en una secuencia contigua
	// de memoria

	// Los arrays en go tienen un tamaño definido

	// Delaración clásica
	var foods [3]string

	foods[0] = "Pizza"
	foods[1] = "Tallarines"
	foods[2] = "Ceviche"

	fmt.Println(foods)

	// Array literal
	numbers := [3]uint8{1, 2, 3}
	fmt.Println(numbers)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
