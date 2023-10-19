package main

import (
	"errors"
	"io"
)

func NewBox() *Box {
	return &Box{
		Width:  1,
		Height: 1,
	}
}

func (b *Box) PrintLineNumber(lineNumber int, w io.Writer) (err error) {
	if b.Height > lineNumber {
		line := Line{
			Content: "",
			Size:    0,
		}

		// if no more lines are available and height is greater than line number, fill up the empty line
		if len(b.Lines) <= lineNumber && lineNumber < b.Height {
			var text string
			text, err = fillUpEmptyLine(b)

			line.Content = text
			line.Size = len(text)

			if err != nil {
				return err
			}
		} else {
			line = b.Lines[lineNumber]

			if line.Size < b.Width {
				for i := line.Size; i < b.Width; i++ {
					line.Content += " "
				}
			}

			err = addBackgroundColor(*b, &line.Content)
			if err != nil {
				return err
			}
		}

		line.Content = AddResetCode(line.Content)

		PrintRaw(line.Content, w)

		return nil
	} else {
		return errors.New(END_OF_BOX_REACHED_ERROR)
	}
}

func (b *Box) AddView(viewContentGenerator ViewContentGenerator) error {
	lines, sizes, err := viewContentGenerator.GetViewContent()
	if err != nil {
		return err
	}
	for i, l := range lines {
		line := Line{
			Content: l,
			Size:    sizes[i],
		}
		b.Lines = append(b.Lines, line)
	}

	return nil
}

func fillUpEmptyLine(box *Box) (string, error) {
	var newLine string
	for i := 0; i < box.Width; i++ {
		newLine += " "
	}

	err := addBackgroundColor(*box, &newLine)
	if err != nil {
		return "", err
	}

	return newLine, nil
}

func addBackgroundColor(box Box, line *string) error {
	if box.BackgroundColor != "" {
		var err error
		*line, err = WrapContentWithBackgroundColor(*line, box.BackgroundColor)
		if err != nil {
			return err
		}
	}
	return nil
}
