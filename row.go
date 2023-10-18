package main

import (
	"io"
)

func NewRow() *Row {
	return &Row{}
}

func (r *Row) Print(w io.Writer) {
	cursorPositionVertical := 0
	numBoxes := len(r.Boxes)
	endOfBoxReachedCounter := 0

	for {
		for _, box := range r.Boxes {
			err := box.PrintLineNumber(cursorPositionVertical, w)
			if err != nil && err.Error() == END_OF_BOX_REACHED_ERROR {
				endOfBoxReachedCounter++
			}
		}
		if numBoxes == endOfBoxReachedCounter {
			break
		}

		PrintRaw("\n", w)

		cursorPositionVertical++
	}
}

func (r *Row) Add(box *Box) {
	r.Boxes = append(r.Boxes, box)
}
