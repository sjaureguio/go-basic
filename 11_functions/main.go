package main

import (
	"fmt"
	"strings"
)

func main() {
	// Función con parametros por valor
	name := "Charles"
	hello(name)

	// Paso de paramétros por referencia
	change(&name)
	fmt.Println(name)

	// Función que retorna mas de un valor
	minus, mayus := convert(name)
	fmt.Println(minus, mayus)

	// Funcion que recibe una función como paramétro
	nums := []int{1, 4, 70, 5, 67, 90, 2}
	result := filter(nums, func(i int) bool {
		if i <= 10 {
			return true
		}

		return false
	})
	fmt.Println(result)

	// Función que retorna una función
	hello := great("Solano")
	fmt.Println(hello("Jauregui"))

	// Función variática
	total := sum(2, 3, 89, 453, 9)
	fmt.Println(total)

	// Función anónima
	func() {
		fmt.Println("Hello Sislii")
	}()
}

// Función con parametros por valor
func hello(name string) {
	fmt.Printf("Hello %s \n", name)
}

// Paso de paramétros por referencia
func change(name *string) {
	*name = "Solano Jauregui"
}

// Función que retorna mas de un valor
func convert(text string) (string, string) {
	minus := strings.ToLower(text)
	mayus := strings.ToUpper(text)

	return minus, mayus
}

// Funcion que recibe una función como paramétro
func filter(nums []int, callback func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}

	return result
}

// Función que retorna una función
func great(name string) func(string) string {
	return func(text string) string {
		return fmt.Sprintf("Hello %s %s", name, text)
	}
}

// Función variática
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
