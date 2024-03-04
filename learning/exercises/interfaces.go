package exercises

import (
	"fmt"

	interfaces "github.com/juparefe/Golang-Ecommerce/learning/interfaces"
)

func HumansBreathing(hu interfaces.Human) {
	hu.Breathe()
	fmt.Printf("Soy un/a %s y estoy respirando \n", hu.Gender())
}
