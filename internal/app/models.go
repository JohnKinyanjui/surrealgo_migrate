package app

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func readModel(model interface{}) {
	var tableName string
	// var fields []string

	val := reflect.ValueOf(model)
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("Field: %s, Tag: %q\n", field.Name, field.Tag)

	}

	if len(tableName) == 0 {
		tableName = convertToSnakeCase(tableName)
	}
}

func convertToSnakeCase(str string) string {
	var result strings.Builder
	var prevChar rune
	for _, char := range str {
		if unicode.IsUpper(char) {
			if prevChar != 0 && unicode.IsLower(prevChar) {
				result.WriteRune('_')
			}
			char = unicode.ToLower(char)
		}
		prevChar = char
		result.WriteRune(char)
	}
	return result.String()
}
