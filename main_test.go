package main

import (
	"os"
	"testing"
)

func TestPrintRaw(t *testing.T) {
	writer := os.Stdout

	PrintRaw("Hello, World!", writer)
	PrintRaw("\nHello, World!\n", writer)
}
