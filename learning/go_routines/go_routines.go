package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func SlowName(name string) {
	letter := strings.Split(name, "")
	for _, letter := range letter {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
