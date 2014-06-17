package netstring

import (
	"fmt"
	"testing"
)

func TestMarshall(t *testing.T) {
	output := Marshall("Hello World!", "cat", "dog")
	if output != "12:Hello World!,3:cat,3:dog," {
		t.Error("marshall output does not match expectation")
	}
}

func TestUnmarshall(t *testing.T) {
	components, err := Unmarshall("12:Hello World!,3:cat,3:dog,")
	expectedResults := []string{"Hello World!", "cat", "dog"}
	if err != nil {
		t.Error("error should be nil")
	}
	for i, expected := range expectedResults {
		if expected != components[i] {
			t.Errorf("Expected %s received %s\n", expected, components[i])
		}
	}
}

func TestErrorChecking(t *testing.T) {
	_, err := Unmarshall("1a")
	if err == nil {
		t.Error("minimum length error should have occurred")
	}
	_, err = Unmarshall("12:hello,")
	if err == nil {
		t.Error("invalid length error should have occurred")
	}
	_, err = Unmarshall("6:hello")
	if err == nil {
		t.Error("missing comma error should have occurred")
	}
	_, err = Unmarshall(":hello")
	if err == nil {
		t.Error("invalid length error should have occurred")
	}
}

func ExampleMarshall() {
	output := Marshall("Hello World!", "cat", "dog")
	fmt.Println(output)
	// Output: 12:Hello World!,3:cat,3:dog,
}

func ExampleUnmarshall() {
	parsedNetstrings, err := Unmarshall("12:Hello World!,3:cat,3:dog,")
	if err != nil {
		fmt.Println(err)
	}
	for _, ns := range parsedNetstrings {
		fmt.Println(ns)
	}
	// Output:
	// Hello World!
	// cat
	// dog

}
