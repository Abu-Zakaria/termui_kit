package main

import "io"

func (tv *TextView) SetContent(content string) {
	tv.content = content
}

func (tv *TextView) GetContent() string {
	return tv.content
}

func (tv *TextView) Print(w io.Writer) error {
	wrapped_content, err := WrapContentWithColor(tv.content, tv.color)
	if err != nil {
		panic(err)
	}

	if !tv.SingleLine {
		wrapped_content += "\n"
	}

	PrintRaw(wrapped_content, w)
	return nil
}

func (tv *TextView) SetColor(color string) error {
	tv.color = color
	return nil
}
