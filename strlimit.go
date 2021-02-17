// Package strlimit provides functions to limit the bytes of a string with rune-safe.
package strlimit

import (
	"strings"
	"unicode/utf8"
)

// LimitBytes limits the number of bytes in a string with rune-safe.
// The bytes of a returned string is a given limit bytes or lower.
func LimitBytes(s string, limit int) string {
	return LimitBytesWithEnd(s, limit, "")
}

// LimitBytesWithEnd limits the number of bytes in a string with rune-safe,
// and append a given end string.
// The bytes of a returned string is a given limit bytes or lower.
func LimitBytesWithEnd(s string, limit int, end string) string {
	var b strings.Builder
	rem := limit - len(end)

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if rem-size < 0 {
			b.WriteString(end)
			break
		}
		rem -= size
		s = s[size:]

		b.WriteRune(r)
	}

	return b.String()
}
