package versions

import "strings"

func trimNewline(str string) string {
	return strings.TrimRight(str, "\r\n")
}
