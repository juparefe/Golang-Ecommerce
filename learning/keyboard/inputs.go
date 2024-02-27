package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func NumbersInput() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el numero 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto => " + err.Error())
		}
	}

	fmt.Println("Ingrese el numero 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto => " + err.Error())
		}
	}

	fmt.Println("Ingrese la leyenda: ")
	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println(legend, "=>", number1*number2)
}
