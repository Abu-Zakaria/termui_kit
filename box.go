package main

import (
	"errors"
	"io"
)

func NewBox() *Box {
	return &Box{
		Width:  1,
		Height: 1,
		Padding: []int{
			0, // top
			0, // right
			0, // bottom
			0, // left
		},
	}
}

func (b *Box) PrintLine(lineNumber int, w io.Writer) (err error) {
	if b.GetRealHeight() > lineNumber {
		line := Line{
			Content: "",
			Size:    0,
		}

		// if no more lines are available and height is greater than line number, fill up the empty line
		if len(b.Lines) <= lineNumber {
			var text string
			var size int

			text, size, err = fillUpEmptyLine(b)

			line.Content = text
			line.Size = size

			if err != nil {
				return err
			}
		} else {
			line = b.Lines[lineNumber]

			if line.Size < b.Width {
				fillUpText := ""
				for i := line.Size; i < b.Width; i++ {
					fillUpText += " "
				}
				fillUpText, err = WrapContentWithBackgroundColor(fillUpText, b.BackgroundColor)
				if err != nil {
					return err
				}
				line.Content += AddResetCode(fillUpText)
			}
		}

		PrintRaw(line.Content, w)

		return nil
	} else {
		return errors.New(END_OF_BOX_REACHED_ERROR)
	}
}

func (b *Box) AddView(viewContentGenerator ViewContentGenerator) (err error) {
	topPaddingLines, err := getTopPaddingLines(*b)
	if err != nil {
		return err
	}

	b.Lines = append(b.Lines, topPaddingLines...)

	lines := []string{}
	sizes := []int{}

	lines, sizes, err = viewContentGenerator.GetViewContent(*b)
	if err != nil {
		return err
	}

	for i, l := range lines {
		line := Line{
			Content: l,
			Size:    sizes[i],
		}
		applyHorizontalPadding(*b, &line)

		b.Lines = append(b.Lines, line)
	}

	bottomPaddingLines, err := getBottomPaddingLines(*b)
	if err != nil {
		return err
	}

	b.Lines = append(b.Lines, bottomPaddingLines...)

	return nil
}

func (b *Box) GetRealHeight() int {
	return b.Height + b.Padding[0] + b.Padding[2]
}

func (b *Box) GetRealWidth() int {
	return b.Width + b.Padding[1] + b.Padding[3]
}

func fillUpEmptyLine(box *Box) (string, int, error) {
	var newLine string
	var size int

	for i := 0; i < box.GetRealWidth(); i++ {
		newLine += " "
		size++
	}

	err := addBackgroundColor(*box, &newLine)
	if err != nil {
		return "", 0, err
	}

	return AddResetCode(newLine), size, nil
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

func getTopPaddingLines(b Box) (topPaddingLines []Line, err error) {
	for i := 0; i < b.Padding[0]; i++ {
		paddingLine := NewLine()

		err = addBackgroundColor(b, &paddingLine.Content)
		if err != nil {
			return []Line{}, err
		}

		for j := 0; j < b.GetRealWidth(); j++ {
			paddingLine.Content += " "
			paddingLine.Size++
		}

		paddingLine.Content = AddResetCode(paddingLine.Content)

		topPaddingLines = append(topPaddingLines, *paddingLine)
	}

	return topPaddingLines, nil
}

func getBottomPaddingLines(b Box) (bottomPaddingLines []Line, err error) {
	for i := 0; i < b.Padding[2]; i++ {
		paddingLine := NewLine()

		err = addBackgroundColor(b, &paddingLine.Content)
		if err != nil {
			return []Line{}, err
		}

		for j := 0; j < b.GetRealWidth(); j++ {
			paddingLine.Content += " "
			paddingLine.Size++
		}

		paddingLine.Content = AddResetCode(paddingLine.Content)

		bottomPaddingLines = append(bottomPaddingLines, *paddingLine)
	}

	return bottomPaddingLines, nil
}

func applyHorizontalPadding(box Box, line *Line) (err error) {
	var rightPadding string
	var leftPadding string

	for i := 0; i < box.Padding[1]; i++ {
		rightPadding += " "
	}
	for i := 0; i < box.Padding[3]; i++ {
		leftPadding += " "
	}

	leftPadding, err = WrapContentWithBackgroundColor(leftPadding, box.BackgroundColor)
	if err != nil {
		return err
	}

	rightPadding, err = WrapContentWithBackgroundColor(rightPadding, box.BackgroundColor)
	if err != nil {
		return err
	}

	line.Content = AddResetCode(leftPadding) + line.Content + AddResetCode(rightPadding)

	return nil
}
