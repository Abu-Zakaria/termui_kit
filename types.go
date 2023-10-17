package main

type Kit struct {
	Elements []Element
}

type Element interface {
	SetContent(content string)

	GetContent() string
}

type TextView struct {
	content string
}
