package exercises

import (
	"fmt"
	"strconv"
)

func ConvertToInteger(text string) {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Hubo un error" + err.Error())
	}
	var hundred string
	if number > 1000 {
		hundred = "Es mayor a 100"
	} else {
		hundred = "Es menor a 100"
	}
	if number > 10 {
		fmt.Println("Exercise01:", number)
		fmt.Println("Exercise01:", hundred)
	}
}
