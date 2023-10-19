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
	box4 := NewBox()
	box5 := NewBox()

	box1.Width, box2.Width, box3.Width, box4.Width, box5.Width = 20, 30, 25, 40, 25
	box1.Height, box2.Height, box3.Height, box4.Height = 3, 7, 2, 4
	box1.BackgroundColor, box2.BackgroundColor, box3.BackgroundColor, box4.BackgroundColor, box5.BackgroundColor = "red", "green", "blue", "magenta", "blue"
	box1.Padding[0] = 2
	box1.Padding[1] = 3
	box1.Padding[2] = 4
	box1.Padding[3] = 8

	box3.Padding[0] = 10
	box3.Padding[1] = 3
	box3.Padding[2] = 20
	box3.Padding[3] = 15

	box4.Padding[0] = 2

	textView := NewTextView()

	textView.Content = "Hello World!"
	textView.Color = "green"

	box1.AddView(textView)

	tv2 := NewTextView()

	tv2.Content = "This is second sentence."
	tv2.Color = "blue"

	box2.AddView(tv2)

	tv3 := NewTextView()

	tv3.Content = "This is third sentence."
	tv3.Color = "red"

	box3.AddView(tv3)

	tv4 := NewTextView()

	tv4.Content = "This is fourth sentence."
	tv4.Color = "red"

	box4.AddView(tv4)

	tv5 := NewTextView()

	tv5.Content = "This is fifth sentence."
	tv5.Color = "green"

	box5.AddView(tv5)

	row.Add(box1)
	row.Add(box2)
	row.Add(box3)
	row.Add(box4)
	row.Add(box5)

	kit.Print()
}
