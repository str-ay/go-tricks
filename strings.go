package tricks

import (
	"unicode"
	"unicode/utf8"
)

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent(delimiter byte) {
	if len(b.r) > 0 {
		b.r = append(b.r, delimiter)
	}
}

// CamelCaseToLowerWithDelimiter convert ThisText to this-text or this_text
func CamelCaseToLowerWithDelimiter(s string, delimiter byte) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent(delimiter)
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent(delimiter)
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent(delimiter)
		}
		b.write(m)
	}
	return string(b.r)
}
