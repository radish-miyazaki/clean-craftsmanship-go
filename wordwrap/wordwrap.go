package wordwrap

import "strings"

func Wrap(s string, w int) string {
	if w >= len(s) {
		return s
	}

	li := strings.LastIndex(s[:w+1], " ") + 1
	if li == 0 {
		li = w
	}

	return strings.Trim(s[:li], " ") + "\n" + Wrap(strings.Trim(s[li:], " "), w)
}
