package defer_panic

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el segundo mensaje")
}

func ShowPanic() {
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}

func ShowRecover() {
	defer func() {
		recover := recover()
		if recover != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", recover)
		}
	}()
	a := 2
	if a == 2 {
		panic("Se encontro el valor 2")
	}
}
