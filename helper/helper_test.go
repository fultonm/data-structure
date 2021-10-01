package helper

import (
	"testing"
)

func TestContainsString(t *testing.T) {
	test1 := "abc"
	char1 := rune('s')
	if ContainsString(test1, char1) > 0 {
		t.Fatalf("Did not expected string: %v to have element %v\n", test1, string(char1))
	}
	char2 := rune('c')
	if ContainsString(test1, char2) < 0 {
		t.Fatalf("Expected string: %v to have element %v\n", test1, string(char2))

	}
}

func TestContainsRune(t *testing.T) {
	test1 := []rune("abc")
	if ContainsRune(test1, rune('c')) < 0 {
		t.Fatalf("Expected rune array: %v to have element %v\n", test1, rune('c'))
	}
	if ContainsRune(test1, rune('d')) > 0 {
		t.Fatalf("Did not expect rune array: %v to have element %v\n", test1, rune('d'))
	}
}
