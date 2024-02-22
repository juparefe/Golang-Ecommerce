package variables

import (
	"strconv"
)

func ConvertToInteger(text string) (int, string) {
	number, _ := strconv.Atoi(text)
	var hundred string
	if number > 100 {
		hundred = "Es mayor a 100"
	} else {
		hundred = "Es menor a 100"
	}
	return number, hundred
}
