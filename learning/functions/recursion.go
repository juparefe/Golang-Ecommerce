package functions

import (
	"fmt"
)

// La funcion se ejecuta hasta que se cumpla la condicion que asigne en el if
func Exponentiation(value int) {
	if value > 100000 {
		return
	}
	fmt.Println(value)
	Exponentiation(value * 2)
}
