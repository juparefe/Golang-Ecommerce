package functions

import "fmt"

func returnFunction(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	multiplicationTable := 2
	MyMultiplicationTable := returnFunction(multiplicationTable)
	for i := 1; i <= 10; i++ {
		fmt.Println(MyMultiplicationTable())
	}
}
