package protocol

import "strings"

func reverse(s string) string {
	var str strings.Builder
	str.Grow(len(s) - 1)
	for idx := range s {
		str.Write(([]byte{s[(len(s)-1)-idx]}))
	}
	return str.String()
}

func bytesToStr(b []byte) string {
	return reverse(strings.TrimRight(string(b), "\x00"))
}
