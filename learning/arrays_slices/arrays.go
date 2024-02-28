package arrays_slices

import (
	"fmt"
)

var table [10]int = [10]int{10, 9, 8}
var matrix [20][30]int

func ShowArrays() {
	table[7] = 33
	table[3] = 10
	stringTable := [10]string{"Pedro", "Juan"}
	fmt.Println(table)
	for i := 0; i < len(stringTable); i++ {
		fmt.Println(stringTable[i])
	}
	matrix[7][4] = 15
	fmt.Println(matrix)
}
