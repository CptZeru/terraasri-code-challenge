package utils

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

// ThrowTestingErr used to throw err message for *test.go
func ThrowTestingErr(t *testing.T, expect interface{}, result interface{}) {
	t.Errorf("Expected: %v, Result: %v", expect, result)
}

// readStringInput acts as reader for user input as a string.
// It will remove the \n at the end of every input.
func ReadStringInput(reader *bufio.Reader, inputMessage string) string {
	if inputMessage == "" {
		fmt.Print("Enter input: ")
	} else {
		fmt.Printf("%v", inputMessage)
	}
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}
