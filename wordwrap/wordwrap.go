package wordwrap

import "strings"

func Wrap(s string, w int) string {
	//return strings.ReplaceAll(s, " ", "\n")
	if w >= len(s) {
		return s
	}

	li := strings.LastIndex(s[:w+1], " ") + 1
	if li == 0 {
		li = w
	}

	return strings.Trim(s[:li], " ") + "\n" + Wrap(strings.Trim(s[li:], " "), w)
}

// "xx xx" -> "xx" + "xx"

// "x x x" -> "x x\n" + "x"
