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
		Writer: WRITER,
	}
}

func (kit *Kit) AddRow(row *Row) {
	kit.Rows = append(kit.Rows, row)
}

func (kit *Kit) Print() {
	for _, row := range kit.Rows {
		row.Print(kit.Writer)
	}
}
