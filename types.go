package main

import "io"

type Kit struct {
	Rows   []*Row
	Writer io.Writer
}

type Row struct {
	Boxes           []*Box
	BackgroundColor string
	BorderColor     string
}

type Box struct {
	Lines           []Line
	Width           int
	Height          int
	BackgroundColor string
	BorderColor     string
	Padding         []int
}

type Line struct {
	Content string
	Size    int
}

type TextView struct {
	Content             string
	Color               string
	MultiLine           bool
	UseParentBackground bool
}

// 1st return: slice of the lines of the content after adding color
// 2nd return: slice of each line's content size, without counting the color codes
// 3rd return: error if fail to add color to the content
type ViewContentGenerator interface {
	GetViewContent(box Box) ([]string, []int, error)
}
