package main

import (
	"testing"
)

/* func TestPrintRaw(t *testing.T) {
	writer := os.Stdout

	PrintRaw("Hello, World!", writer)
	PrintRaw("\nGreetings, World!", writer)

	// Clears the top line from start to end
	// PrintRaw("\033[1A\033[K", writer)

	PrintRaw("Bye, World!\n", writer)
} */

func TestKit(t *testing.T) {
	kit := NewKit()

	textView := &TextView{}

	textView.SetContent("Hello, World!")

	kit.AddElement(textView)

	kit.Print()
}
