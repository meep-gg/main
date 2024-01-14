package log

import (
	"fmt"
)

func Highlight(color, input interface{}) string {
	if color == "r" {
		// return fmt.Sprintf("\033[91m%v\033[0m", input)
		return fmt.Sprintf("\033[31m%v\033[0m", input)
	} else if color == "y" {
		return fmt.Sprintf("\033[33m%v\033[0m", input)
	} else if color == "b" {
		return fmt.Sprintf("\033[34m%v\033[0m", input)
	} else if color == "g" {
		return fmt.Sprintf("\033[32m%v\033[0m", input)
	} else if color == "p" {
		return fmt.Sprintf("\033[35m%v\033[0m", input)
	} else {
		return fmt.Sprintf("\033[37m%v\033[0m", input)
	}
}

func Info(input interface{}) string {
	return Highlight("b", input)
}

func Error(input interface{}) string {
	return Highlight("r", input)
}

func Warn(input interface{}) string {
	return Highlight("y", input)
}

func Success(input interface{}) string {
	return Highlight("g", input)
}

func Title(input interface{}) string {
	return Highlight("p", input)
}
