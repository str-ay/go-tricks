package tricks

import (
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

// RepeatSeparated repeats string with delimiter
// RepeatSeparated("?", ",", 4) -> "?,?,?,?"
func RepeatSeparated(s string, sep string, number int) string {
	if number < 1 {
		return ""
	}

	return s + strings.Repeat(sep+s, number-1)
}
