package main

import "io"

func (tv *TextView) SetContent(content string) {
	tv.content = content + "\n"
}

func (tv *TextView) GetContent() string {
	return tv.content
}

func (tv *TextView) Print(w io.Writer) error {
	PrintRaw(tv.content, w)
	return nil
}
