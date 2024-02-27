package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error
var textToSave string

func ValidateMistake() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el numero: ")
	if scanner.Scan() {
		var text = scanner.Text()
		num, err = strconv.Atoi(text)
		if err != nil {
			fmt.Println("El numero ingresado (", text, ") no es correcto")
			ValidateMistake()
		}
	}
	for i := 0; i <= 10; i++ {
		fmt.Printf("Exercise02: %d x %d = %d \n", num, i, (num * i))
		textToSave += fmt.Sprintf("%d x %d = %d \n", num, i, (num * i))
	}
	return textToSave
}
