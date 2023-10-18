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

	textView := &TextView{
		SingleLine: true,
	}

	textView.SetContent("This color is red.")
	textView.SetColor("red")

	kit.AddElement(textView)

	textView2 := &TextView{
		SingleLine: true,
	}

	textView2.SetContent("This color is blue.")
	textView2.SetColor("blue")

	kit.AddElement(textView2)

	textView3 := &TextView{}

	textView3.SetContent("This color is green.")
	textView3.SetColor("green")

	kit.AddElement(textView3)

	kit.Print()
}
