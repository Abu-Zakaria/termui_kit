package main

import (
	"io"
	"math"
	"slices"
)

func NewRow() *Row {
	return &Row{}
}

func (r *Row) Print(w io.Writer) {
	cursorPositionVertical := 0
	numBoxes := len(r.Boxes)
	endOfBoxReachedCounter := 0
	fullyPrintedBoxes := []*Box{}
	maxHeight := 0

	for _, box := range r.Boxes {
		maxHeight = int(math.Max(float64(maxHeight), float64(box.Height)))
	}

	for {
		for _, box := range r.Boxes {
			if box.Height < maxHeight && cursorPositionVertical >= box.Height {
				fillUpRowText := fillUpRow(*box)

				PrintRaw(fillUpRowText, w)
			}

			err := box.PrintLineNumber(cursorPositionVertical, w)
			if err != nil && err.Error() == END_OF_BOX_REACHED_ERROR && !slices.Contains(fullyPrintedBoxes, box) {
				endOfBoxReachedCounter++
				fullyPrintedBoxes = append(fullyPrintedBoxes, box)
			}
		}
		if numBoxes <= endOfBoxReachedCounter {
			break
		}

		PrintRaw("\n", w)

		cursorPositionVertical++
	}
}

func (r *Row) Add(box *Box) {
	r.Boxes = append(r.Boxes, box)
}

func fillUpRow(box Box) string {
	var emptyContent string

	for i := 0; i < box.Width; i++ {
		emptyContent += " "
	}

	return emptyContent
}
