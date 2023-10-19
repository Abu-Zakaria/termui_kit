package main

import (
	"errors"
	"io"
)

func NewTextView() *TextView {
	return &TextView{}
}

func (tv *TextView) Print(w io.Writer) error {
	viewContents, _, err := tv.GetViewContent()
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

func (tv *TextView) GetViewContent() ([]string, []int, error) {
	wrapped_content, err := WrapContentWithColor(tv.Content, tv.Color)
	if err != nil {
		return []string{""}, []int{0}, err
	}

	return []string{wrapped_content}, []int{len(tv.Content)}, nil
}
