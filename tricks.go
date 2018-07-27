package tricks

import (
	"bytes"
	"strconv"
)

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
	for i := start; i <= start-1+number; i++ {
		result.WriteString("$")
		result.WriteString(strconv.Itoa(i))
		if i != start-1+number {
			result.WriteString(",")
		}
	}

	return result.String()
}
