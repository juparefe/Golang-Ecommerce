package tools

import (
	"fmt"
	"strings"
	"time"
)

func DateMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func EscapeString(text string) string {
	desc := strings.ReplaceAll(text, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}
