package tools

import "strings"

func ToBright(s string) string {
	n := []rune(s)
	nm := make([]rune, len(n)*2)
	copy(nm, n)
	for i := len(n) - 1; i >= 0; i-- {
		copy(nm[i+1:], nm[i:])
		nm[i] = '`'
	}
	return string(nm)
}

func BrightLen(s string) int {
	s = strings.ReplaceAll(s, "`", "")
	return len([]rune(s))
}
