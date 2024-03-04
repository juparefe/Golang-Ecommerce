package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func AsynchronousName(name string) {
	letter := strings.Split(name, "")
	for _, letter := range letter {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}

func ShowChannels(name string, channel chan bool) {
	letter := strings.Split(name, "")
	for _, letter := range letter {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
	channel <- true
}
