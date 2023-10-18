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
	kit, row := NewKit(), NewRow()

	kit.AddRow(row)

	box1 := NewBox()
	box2 := NewBox()
	box3 := NewBox()

	box1.Width, box2.Width, box3.Width = 50, 50, 50
	box1.BackgroundColor, box2.BackgroundColor, box3.BackgroundColor = "red", "green", "blue"
	textView := NewTextView()

	textView.Content = "Hello World!"
	textView.Color = "green"
	textView.MultiLine = false

	box1.AddView(textView)

	tv2 := NewTextView()

	tv2.Content = "This is second sentence."
	tv2.Color = "blue"

	box2.AddView(tv2)

	tv3 := NewTextView()

	tv3.Content = "This is third sentence."
	tv3.Color = "red"

	box3.AddView(tv3)

	row.Add(box1)
	row.Add(box2)
	row.Add(box3)

	kit.Print()
}
