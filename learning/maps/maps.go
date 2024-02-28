package maps

import (
	"fmt"
)

func ShowMaps() {
	countries := make(map[string]string)
	fmt.Println(countries)
	countries["Mexico"] = "D.F."
	countries["Argentina"] = "Buenos Aires"
	fmt.Println(countries)
	fmt.Println(countries["Argentina"])
	championship := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	fmt.Println(championship)
	for team, score := range championship {
		fmt.Printf("Equipo: %s, tiene un puntaje de: %d \n", team, score)
	}
	delete(championship, "Real Madrid")
	fmt.Println(championship)

	score, exist := championship["Juventus"]
	fmt.Printf("El puntaje es: %d y es %t que el equipo existe", score, exist)
}
