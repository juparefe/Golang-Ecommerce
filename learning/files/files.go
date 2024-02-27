package files

import (
	"bufio"
	"fmt"
	"os"

	exercises "github.com/juparefe/Golang-Ecommerce/learning/exercises"
)

func SaveFile() {
	var textToSave = exercises.ValidateMistake()
	file, err := os.Create("./learning/files/txt/newFile.txt")
	if err != nil {
		fmt.Println("Error al crear archivo " + err.Error())
		return
	}
	fmt.Fprintln(file, textToSave)
	file.Close()
}

func AppendFile() {
	var textToAppend = exercises.ValidateMistake()
	if !Append("./learning/files/txt/updatedFile.txt", textToAppend) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(path string, textToAppend string) bool {
	arch, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}
	_, err = arch.WriteString(textToAppend)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}
	arch.Close()
	return true
}

func ReadFile() {
	file, err := os.Open("./learning/files/txt/updatedFile.txt")
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		registry := scanner.Text()
		fmt.Println("=> " + registry)
	}
	file.Close()
}
