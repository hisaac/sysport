package utils

import "strings"

func TrimNewline(str string) string {
	return strings.TrimRight(str, "\r\n")
}
