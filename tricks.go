package tricks

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
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

// QuestionMarks is useful for dynamic SQL:
// in(" + QuestionMarks(len(ids)) + ") -> in(?,?,?,?,?)
func QuestionMarks(number int) string {
	if number < 1 {
		return ""
	}

	if number == 1 {
		return "?"
	}

	return RepeatSeparated("?", ",", number)
}

// QueryArguments is useful for dynamic SQL:
// in(" + QueryArguments(len(ids)) + ") -> in($1,$2,$3,$4,$5)
func QueryArguments(number, start int) string {
	if number < 1 {
		return ""
	}

	if number == 1 {
		return "$" + strconv.Itoa(start)
	}

	result := bytes.Buffer{}
	for i := start; i <= start+number; i++ {
		result.WriteString("$")
		result.WriteString(strconv.Itoa(i))
		if i != start+number {
			result.WriteString(",")
		}
	}

	return result.String()
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
