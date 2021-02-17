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

// LimitRunes limits the number of runes in a string.
// The number of runes of a returned string is a given limit or lower.
func LimitRunes(s string, limit int) string {
	return LimitRunesWithEnd(s, limit, "")
}

// LimitRunesWithEnd limits the number of runes in a string,
// and append a given end string.
// The number of runes of a returned string is a given limit or lower.
func LimitRunesWithEnd(s string, limit int, end string) string {
	strRunes := utf8.RuneCountInString(s)
	if strRunes <= limit {
		return s
	}

	rem := limit - utf8.RuneCountInString(end)

	var b strings.Builder
	for _, r := range s {
		if rem == 0 {
			break
		}
		b.WriteRune(r)
		rem--
	}
	b.WriteString(end)
	return b.String()
}
