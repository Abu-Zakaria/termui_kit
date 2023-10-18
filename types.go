package main

import "io"

type Kit struct {
	Elements []Element
	Writer   io.Writer
}

type Element interface {
	SetContent(content string)

	GetContent() string

	Print(w io.Writer) error
}

type TextView struct {
	content string
	color   string

	SingleLine bool
}
