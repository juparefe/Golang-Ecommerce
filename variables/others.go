package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func OtherVariables() {
	Name = "Pedro"
	Status = true
	Salary = 1234
	Date = time.Now()

	fmt.Println("Name:", Name)
	fmt.Println("Status:", Status)
	fmt.Println("Salary:", Salary)
	fmt.Println("Date:", Date)
}

func ConvertToText(number int) (bool, string) {
	var text string = strconv.Itoa(number)
	return true, text
}
