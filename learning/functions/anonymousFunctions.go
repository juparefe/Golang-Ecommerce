package functions

import (
	"fmt"
)

func Calculations() {
	var number3 int = 50
	sum := func(number1 int, number2 int) int {
		return number1 + number2 + number3
	}
	fmt.Println(sum(10, 25))
}
