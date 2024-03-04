package main

import (
	"fmt"

	arrays_slices "github.com/juparefe/Golang-Ecommerce/learning/arrays_slices"
	defer_panic "github.com/juparefe/Golang-Ecommerce/learning/defer_panic"
	exercises "github.com/juparefe/Golang-Ecommerce/learning/exercises"
	files "github.com/juparefe/Golang-Ecommerce/learning/files"
	functions "github.com/juparefe/Golang-Ecommerce/learning/functions"
	iterations "github.com/juparefe/Golang-Ecommerce/learning/iterations"
	keyboard "github.com/juparefe/Golang-Ecommerce/learning/keyboard"
	maps "github.com/juparefe/Golang-Ecommerce/learning/maps"
	models "github.com/juparefe/Golang-Ecommerce/learning/models"
	users "github.com/juparefe/Golang-Ecommerce/learning/users"
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
		arrays_slices.ShowArrays()
		arrays_slices.ShowSlices()
		arrays_slices.StorageCapacity()
		maps.ShowMaps()
		users.AddUser()
		Juan := new(models.Man)
		exercises.HumansBreathing(Juan)
		Lina := new(models.Women)
		exercises.HumansBreathing(Lina)
		defer_panic.ShowDefer()
		defer_panic.ShowPanic()
		defer_panic.ShowRecover()
	}

}
