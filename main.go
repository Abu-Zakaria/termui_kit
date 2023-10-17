package main

import (
	"fmt"
	"io"
)

func PrintRaw(text string, w io.Writer) {
	fmt.Fprint(w, text)
}
