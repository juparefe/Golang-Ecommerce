package main

import (
	"fmt"

	exercises "github.com/juparefe/Golang-Ecommerce/learning/exercises"
	files "github.com/juparefe/Golang-Ecommerce/learning/files"
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
	}
	files.ReadFile()
}
