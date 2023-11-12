package main

import "fmt"

func main() {
	// es una estructura de datos que proporciona una colección no ordenada
	// de pares clave-valor, donde cada clave es única. Los maps son implementados
	// como tablas hash, lo que permite un acceso rápido a los valores asociados
	// a través de sus claves

	// Declaración mediante make
	resp := make(map[string]any)

	resp["success"] = true
	resp["message"] = "OK"
	resp["code"] = 200

	// Declaración como map literal
	data := map[string]any{
		"name": "Solano",
		"age":  25,
	}

	// Si una llave del map existe, el valor que retorna
	// es el valor cero del tipo de dato

	if _, ok := data["last_name"]; !ok {
		data["last_name"] = "Jauregui"
	}

	fmt.Println(data)
	fmt.Println(data["last_name"])

	// Cuando se itera map con un for range el indice
	// del for será la llave del map
	for i, val := range resp {
		fmt.Printf("%v, %v \n", i, val)
	}

	// Eliminar key/value pairs de un map
	delete(resp, "code")
	fmt.Println(resp)

}
