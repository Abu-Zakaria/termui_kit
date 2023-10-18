package main

import (
	"errors"
	"io"
)

func NewTextView() *TextView {
	return &TextView{}
}

func (tv *TextView) Print(w io.Writer) error {
	viewContents, err := tv.GetViewContent()
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

func (tv *TextView) GetViewContent() ([]string, error) {
	wrapped_content, err := WrapContentWithColor(tv.Content, tv.Color)
	if err != nil {
		return []string{""}, err
	}

	return []string{wrapped_content}, nil
}
