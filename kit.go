package main

import (
	"os"
)

var (
	// default writer for the kit
	WRITER = os.Stdout
)

func NewKit() *Kit {
	return &Kit{
		Elements: []Element{},
		Writer:   WRITER,
	}
}

func (kit *Kit) AddElement(element Element) {
	kit.Elements = append(kit.Elements, element)
}

func (kit *Kit) Print() {
	for _, element := range kit.Elements {
		element.Print(kit.Writer)
	}
}
