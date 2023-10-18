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
	Lines           []string
	Width           int
	Height          int
	BackgroundColor string
	BorderColor     string
}

type TextView struct {
	Content string
	Color   string

	MultiLine bool
}

type ViewContentGenerator interface {
	GetViewContent() ([]string, error)
}
