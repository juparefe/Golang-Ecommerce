package iterations

import (
	"fmt"
)

func Iterate() {
	// Se comporta como un while, mientras no se ejecute el break se queda iterando
	for {
		fmt.Println("Hola mundo")
		break
	}

	// Se comporta como un for, lo puedo escribir de la forma larga o la corta
	// Forma larga sumando de 1 en 1
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// Forma corta sumando de 1 en 1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Forma corta restando de 1 en 1
	for i := 10; i > 5; i-- {
		fmt.Println(i)
	}
	// Forma corta que va de 6 en 6 sumando
	for i := 0; i < 100; i += 6 {
		fmt.Println(i)
	}
	// Forma corta que va de 7 en 7 restando
	for i := 100; i > 10; i -= 7 {
		fmt.Println(i)
	}
	// Forma corta usando el continue para que no muestre el 6
	for i := 10; i > 4; i-- {
		if i == 7 {
			continue
		}
		fmt.Println(i)
	}
}
