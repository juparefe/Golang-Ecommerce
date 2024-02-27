package variables

import "fmt"

func ShowIntegers() {
	var commonInt int
	var int32 int32 = 10
	var int64 int64 = 100

	fmt.Println("Integer comun:", commonInt)
	fmt.Println("Integer de 32:", int32)
	fmt.Println("Integer de 64:", int64)
}
