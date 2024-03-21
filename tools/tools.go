package tools

import (
	"fmt"
	"strconv"
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

func CreateScript(s, fieldName, typeField, ValueS string, ValueN int, ValueF float64) string {
	if (typeField == "S" && len(ValueS) == 0) ||
		(typeField == "F" && ValueF == 0) ||
		(typeField == "N" && ValueN == 0) {
		return s
	}
	if !strings.HasSuffix(s, "SET ") {
		s += ", "
	}

	switch typeField {
	case "S":
		s += fieldName + " = '" + EscapeString(ValueS) + "'"
	case "N":
		s += fieldName + " = " + strconv.Itoa(ValueN)
	case "F":
		s += fieldName + " = " + strconv.FormatFloat(ValueF, 'e', -1, 64)
	}

	fmt.Println("Script,", s)
	return s
}
