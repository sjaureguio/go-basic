package greet

import "fmt"

func English() string {
	return fmt.Sprintf("Hi %s \n", name)
}

func Italian() string {
	return fmt.Sprintf("Ciao %s \n", name)
}
