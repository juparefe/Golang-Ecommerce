package variables

import (
	"strconv"
)

func ConvertToInteger(text string) (int, string) {
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	var hundred string
	if number > 100 {
		hundred = "Es mayor a 100"
	} else {
		hundred = "Es menor a 100"
	}
	return number, hundred
}
