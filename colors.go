package main

import "errors"

var (
	colors = map[string]string{
		"red":   "\033[31m",
		"blue":  "\033[34m",
		"green": "\033[32m",
	}

	reset_code = "\033[0m"
)

const (
	COLOR_NOT_FOUND_ERROR = "Given color not found in the colors list"
)

func generateColor(color_name string) (string, error) {
	if color, ok := colors[color_name]; ok {
		return color, nil
	}
	return "", errors.New(COLOR_NOT_FOUND_ERROR)
}

func WrapContentWithColor(content string, color string) (string, error) {
	color, err := generateColor(color)
	if err != nil {
		return "", err
	}

	return color + content + reset_code, nil
}
