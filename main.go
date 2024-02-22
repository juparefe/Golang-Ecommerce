package main

import (
	"fmt"

	integers "github.com/juparefe/Golang-Ecommerce/variables"
)

func main() {
	integers.ShowIntegers()
	integers.OtherVariables()
	status, text := integers.ConvertToText(1234)
	fmt.Println("Status:", status)
	fmt.Println("Text:", text)
}
