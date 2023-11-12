package main

import "fmt"

func main() {

	// For range permite iterar sobre slices, maps y strings

	nums := []uint8{2, 4, 6, 8}
	for i, val := range nums {
		fmt.Println(i, val)
	}

	sports := map[string]any{
		"name": "Solano Jauregui",
		"age":  32,
	}
	for key, val := range sports {
		fmt.Println(key, val)
	}

	hello := "Hello"
	for _, val := range hello {
		fmt.Println(string(val))
	}
}
