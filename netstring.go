package netstring

import (
	"bytes"
	"fmt"
	"unicode"
)

// Marshall takes a variadic set of strings and returns them in a single netstring formatted string
func Marshall(components ...string) string {
	var buffer bytes.Buffer
	for _, component := range components {
		buffer.WriteString(fmt.Sprintf("%d:%s,", len(component), component))
	}
	return buffer.String()

}

// Unmarshall takes a netstring formatted string, unmarshalls it and returns a slice of strings and an error
func Unmarshall(ns string) ([]string, error) {
	length := len(ns)
	if length < 3 {
		return nil, fmt.Errorf("invalid format")
	}
	position := 0
	var components []string
	var substringLength int
	var currentLetter rune
	for position < length {
		substringLength = 0
		currentLetter = rune(ns[position])
		for unicode.IsNumber(currentLetter) {
			substringLength *= 10
			substringLength += int(currentLetter - '0')
			position++
			currentLetter = rune(ns[position])
		}
		if substringLength == 0 {
			return nil, fmt.Errorf("length must be 1 or more")
		}
		if length < position+substringLength+1 {
			return nil, fmt.Errorf("invalid length %d specified at %d", substringLength, position-1)
		}
		if currentLetter != ':' {
			return nil, fmt.Errorf("missing colon at position %d", position)
		}
		position++
		components = append(components, ns[position:position+substringLength])
		position += substringLength
		if ns[position] != ',' {
			return nil, fmt.Errorf("no comma at position %d", position)
		}
		position++
	}
	return components, nil

}
