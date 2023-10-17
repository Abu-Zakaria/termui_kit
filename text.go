package main

func (tv *TextView) SetContent(content string) {
	tv.content = content
}

func (tv *TextView) GetContent() string {
	return tv.content
}
