package tricks

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"unicode"
	"unicode/utf8"
)

//JoinTrimmed makes Join with trim all blank spaces like [ ], \t, "\n"
func JoinTrimmed(sep string, str ...string) string {
	trimmed := make([]string, 0, len(str))
	for _, str := range str {
		s := strings.TrimSpace(str)
		if s == "" {
			continue
		}

		trimmed = append(trimmed, s)
	}

	return strings.Join(trimmed, sep)
}

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

// RepeatSeparated repeats string with delimiter
// RepeatSeparated("?", ",", 4) -> "?,?,?,?"
func RepeatSeparated(s string, sep string, number int) string {
	if number < 1 {
		return ""
	}

	return s + strings.Repeat(sep+s, number-1)
}

//MD5Hash generates MD5 string
func MD5Hash(text string) (string, error) {
	md5 := md5.New()
	if _, err := md5.Write([]byte(text)); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5.Sum(nil)), nil
}

// UniqueStrings ...
func UniqueStrings(strs []string) []string {
	stringsMap := make(map[string]bool)
	for _, str := range strs {
		stringsMap[str] = true
	}

	results := make([]string, 0, len(stringsMap))
	for k := range stringsMap {
		results = append(results, k)
	}

	return results
}

func TrimToNil(str string) (strPtr *string) {
	trimmed := strings.TrimSpace(str)
	if trimmed == "" {
		return strPtr
	}

	return &str
}
