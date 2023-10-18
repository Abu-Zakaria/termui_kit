package main

import (
	"errors"
	"io"
)

func NewBox() *Box {
	return &Box{}
}

func (b *Box) PrintLineNumber(lineNumber int, w io.Writer) (err error) {
	if len(b.Lines) > lineNumber {
		line := b.Lines[lineNumber]

		if len(line) < b.Width {
			for i := len(line); i < b.Width; i++ {
				line += " "
			}
		}

		if b.BackgroundColor != "" {
			line, err = WrapContentWithBackgroundColor(line, b.BackgroundColor)
			if err != nil {
				return err
			}
		}

		line = AddResetCode(line)

		PrintRaw(line, w)
		return nil
	} else {
		return errors.New(END_OF_BOX_REACHED_ERROR)
	}
}

func (b *Box) AddView(viewContentGenerator ViewContentGenerator) error {
	content, err := viewContentGenerator.GetViewContent()
	if err != nil {
		return err
	}

	b.Lines = append(b.Lines, content...)

	return nil
}
