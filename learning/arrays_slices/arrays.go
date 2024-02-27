package arrays_slices

import (
	"fmt"
)

var table [10]int

func ShowArrays() {
	table[7] = 33
	table[2] = 10
	fmt.Println(table)
}
