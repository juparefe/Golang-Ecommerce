package main

import (
	"fmt"

	arrays_slices "github.com/juparefe/Golang-Ecommerce/learning/arrays_slices"
	defer_panic "github.com/juparefe/Golang-Ecommerce/learning/defer_panic"
	exercises "github.com/juparefe/Golang-Ecommerce/learning/exercises"
	files "github.com/juparefe/Golang-Ecommerce/learning/files"
	functions "github.com/juparefe/Golang-Ecommerce/learning/functions"
	goroutines "github.com/juparefe/Golang-Ecommerce/learning/go_routines"
	iterations "github.com/juparefe/Golang-Ecommerce/learning/iterations"
	keyboard "github.com/juparefe/Golang-Ecommerce/learning/keyboard"
	maps "github.com/juparefe/Golang-Ecommerce/learning/maps"
	middlewares "github.com/juparefe/Golang-Ecommerce/learning/middlewares"
	models "github.com/juparefe/Golang-Ecommerce/learning/models"
	users "github.com/juparefe/Golang-Ecommerce/learning/users"
	variables "github.com/juparefe/Golang-Ecommerce/learning/variables"
	webserver "github.com/juparefe/Golang-Ecommerce/learning/webserver"
)

func main() {
	var showLearning = false
	if showLearning {
		// Tipos de variables
		variables.ShowIntegers()
		variables.OtherVariables()
		status, text := variables.ConvertToText(1234)
		fmt.Println("Correct:", status)
		fmt.Println("Text:", text)
		variables.CheckSystem()
		// Recibir datos del usuario
		keyboard.NumbersInput()
		// Ciclos
		iterations.Iterate()
		// Manejo de archivos
		files.SaveFile()
		files.AppendFile()
		files.ReadFile()
		// Funciones
		functions.Calculations()
		functions.CallClosure()
		functions.Exponentiation(2)
		// Arrays, slices y maps
		arrays_slices.ShowArrays()
		arrays_slices.ShowSlices()
		arrays_slices.StorageCapacity()
		maps.ShowMaps()
		// Interfaces y modelos
		users.AddUser()
		Juan := new(models.Man)
		exercises.HumansBreathing(Juan)
		Lina := new(models.Women)
		exercises.HumansBreathing(Lina)
		// Defer, panic y recover
		defer_panic.ShowDefer()
		defer_panic.ShowPanic()
		defer_panic.ShowRecover()
		// Asincronia
		goroutines.AsynchronousName("Juan Test")
		go goroutines.AsynchronousName("Juan Test")
		fmt.Println("Estoy aqui")
		var x string
		fmt.Scanln(&x)
		// Canales
		channel := make(chan bool)
		go goroutines.ShowChannels("Juan Test", channel)
		defer func() {
			<-channel
		}()
		// Manejo de servidores
		webserver.MyWebServer()
		// Middlewares
		// Ejercicios
		exercises.ConvertToInteger("500")
		fmt.Println(exercises.ValidateMistake())
	}
	middlewares.MyMiddleware()
}
