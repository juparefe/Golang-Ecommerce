package arrays_slices

import (
	"fmt"
)

var sliceTable []int = []int{10, 9, 8}
var array [10]int = [10]int{10, 9, 8, 7, 6, 5, 4}

func ShowSlices() {
	fmt.Println(sliceTable)

	slice := array[4:]
	slice2 := array[:4]
	slice3 := array[4:5]
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func StorageCapacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo: %d, Capacidad: %d", len(nums), cap(nums))
}
