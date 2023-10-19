package main

import (
	"errors"
	"io"
)

func NewTextView() *TextView {
	return &TextView{
		UseParentBackground: true,
	}
}

func (tv *TextView) Print(box Box, w io.Writer) error {
	viewContents, _, err := tv.GetViewContent(box)
	if err != nil {
		return err
	}

	if len(viewContents) == 0 {
		return errors.New(EMPTY_CONTENT_ERROR)
	}

	content := viewContents[0]

	if tv.MultiLine {
		content += "\n"
	}

	PrintRaw(content, w)

	return nil
}

func (tv *TextView) GetViewContent(box Box) ([]string, []int, error) {
	wrapped_content, err := WrapContentWithColor(tv.Content, tv.Color)
	if err != nil {
		return []string{""}, []int{0}, err
	}
	wrapped_content, err = WrapContentWithBackgroundColor(wrapped_content, box.BackgroundColor)
	if err != nil {
		return []string{""}, []int{0}, err
	}
	wrapped_content = AddResetCode(wrapped_content)

	return []string{wrapped_content}, []int{len(tv.Content)}, nil
}
