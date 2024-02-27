package variables

import (
	"fmt"
	"runtime"
)

func CheckSystem() {
	var os = runtime.GOOS
	if os == "Linux." {
		fmt.Println("If simple: Si es Linux")
	} else {
		fmt.Println("If simple: No es Linux")
	}
	if arq := runtime.GOARCH; arq == "amd64" && (os == "Linux." || os == "windows") {
		fmt.Println("If compuesto: Si es amd64 y Linux-Windows")
	} else {
		fmt.Println("If compuesto: No es amd64 y Linux-Windows")
	}
	switch root := runtime.GOROOT(); root {
	case "go":
		fmt.Println("Switch: El root es la carpeta go")
	default:
		fmt.Println("Switch: El root es la carpeta", root)
	}
}
