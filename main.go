package main

import (
	"fmt"

	arrays_slices "github.com/juparefe/Golang-Ecommerce/learning/arrays_slices"
	exercises "github.com/juparefe/Golang-Ecommerce/learning/exercises"
	files "github.com/juparefe/Golang-Ecommerce/learning/files"
	functions "github.com/juparefe/Golang-Ecommerce/learning/functions"
	iterations "github.com/juparefe/Golang-Ecommerce/learning/iterations"
	keyboard "github.com/juparefe/Golang-Ecommerce/learning/keyboard"
	variables "github.com/juparefe/Golang-Ecommerce/learning/variables"
)

func main() {
	var showLearning = false
	if showLearning {
		variables.ShowIntegers()
		variables.OtherVariables()
		status, text := variables.ConvertToText(1234)
		fmt.Println("Correct:", status)
		fmt.Println("Text:", text)
		variables.CheckSystem()
		exercises.ConvertToInteger("500")
		keyboard.NumbersInput()
		iterations.Iterate()
		fmt.Println(exercises.ValidateMistake())
		files.SaveFile()
		files.AppendFile()
		files.ReadFile()
		functions.Calculations()
		functions.CallClosure()
		functions.Exponentiation(2)
	}
	arrays_slices.ShowArrays()
}
