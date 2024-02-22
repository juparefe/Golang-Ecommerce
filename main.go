package main

import (
	"fmt"
	"runtime"

	exercises "github.com/juparefe/Golang-Ecommerce/exercises"
	integers "github.com/juparefe/Golang-Ecommerce/variables"
)

func main() {
	integers.ShowIntegers()
	integers.OtherVariables()
	status, text := integers.ConvertToText(1234)
	fmt.Println("Correct:", status)
	fmt.Println("Text:", text)
	os := runtime.GOOS
	if os == "Linux." {
		fmt.Println("If simple: Si es Linux")
	} else {
		fmt.Println("If simple: No es Linux")
	}
	if arq := runtime.GOARCH; arq == "amd64" && (os == "Linux." || os == "windows") {
		fmt.Println("If compuesto: Si es amd64 y Linux-Windows")
	} else {
		fmt.Println("If compuesto: No es amd64 y Linux-Windows")
	}
	switch root := runtime.GOROOT(); root {
	case "go":
		fmt.Println("Switch: El root es la carpeta go")
	default:
		fmt.Println("Switch: El root es la carpeta", root)
	}
	number, hundred := exercises.ConvertToInteger("500")
	fmt.Println("Exercise01:", number)
	fmt.Println("Exercise01:", hundred)
}
