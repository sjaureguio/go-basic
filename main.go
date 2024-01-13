package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var name = "Solano"
	fmt.Println(name)

	lastName := "Jauregui"
	fmt.Println(lastName)

	eagle, parrot := "Aguila", "Loro"
	fmt.Println(eagle, parrot)

	switch eagle {
	case "Monkey":
		fmt.Println("Es un burro")
	case "Parrot":
		fmt.Println("Es un loro")
	case "Aguila":
		fmt.Println("Es una águila")
	default:
		fmt.Println("NO es un animal")
	}

	nums := []int{1, 2, 3, 5, 5, 6, 9, 89}

	for i, v := range nums {
		fmt.Println(i, v)
	}

	resp := map[string]any{
		"success": true,
		"message": "OK!",
		"data":    []string{"A", "B", "C", "D"},
	}
	fmt.Printf("P: %p, V: %+v \n", resp, resp)

	steps := [3]string{"Basic", "Plus", "Advanced"}
	for i, v := range steps {
		fmt.Println(i, v)
	}

	i := 0
	for i < 10 {
		fmt.Printf("Index: %v \n", i)
		i++
	}

	//type User struct {
	//	ID uint
	//	Name string
	//	Password string
	//}

	total := sum(1, 3, 6, 8, 6, 9, 9)
	fmt.Println(total)

	fmt.Println(filter(nums, func(num int) bool {
		if num < 5 {
			return true
		}
		return false
	}))
}

// Función variática
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func filter(nums []int, callback func(num int) bool) []int {
	var result []int

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
