package middlewares

import "fmt"

func sum(a, b int) int {
	return a + b
}

func rest(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Inicio")
	result := operationsMiddleware(sum)(2, 3)
	fmt.Println(result)
	result = operationsMiddleware(rest)(12, 3)
	fmt.Println(result)
	result = operationsMiddleware(multiply)(2, 3)
	fmt.Println(result)
}

func operationsMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion:", a, "y", b)
		return f(a, b)
	}
}
