package main

import (
	"errors"
)

var (
	colors = map[string]string{
		"red":   "\033[31m",
		"blue":  "\033[34m",
		"green": "\033[32m",
	}

	backgroundColors = map[string]string{
		"red":   "\033[41m",
		"green": "\033[42m",
		"blue":  "\033[44m",
	}

	resetCode = "\033[0m"
)

const (
	typeColor = iota
	typeBackgroundColor
)

func generateColor(color_name string, t int) (string, error) {
	switch t {
	case typeColor:
		if color, ok := colors[color_name]; ok {
			return color, nil
		}
	case typeBackgroundColor:
		if backgroundColor, ok := backgroundColors[color_name]; ok {
			return backgroundColor, nil
		}
	}
	return "", errors.New(COLOR_NOT_FOUND_ERROR)
}

func WrapContentWithColor(content string, color string) (string, error) {
	color, err := generateColor(color, typeColor)
	if err != nil {
		return "", err
	}

	return color + content, nil
}

func WrapContentWithBackgroundColor(content string, color string) (string, error) {
	backgroundColor, err := generateColor(color, typeBackgroundColor)
	if err != nil {
		return "", err
	}

	return backgroundColor + content, nil
}

func AddResetCode(content string) string {
	return content + resetCode
}
