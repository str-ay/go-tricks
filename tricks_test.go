package tricks

import (
	"testing"
)

func TestGenerateQuestionMarks(t *testing.T) {
	if QuestionMarks(10) != "?,?,?,?,?,?,?,?,?,?" {
		panic("at the disco")
	}
}

func TestQueryArguments(t *testing.T) {
	arr := []string{"1", "2", "3", "4"}
	if QueryArguments(len(arr), 3) != "$3,$4,$5,$6,$7" {
		panic("at the disco")
	}

	arr = []string{"2"}
	if result := QueryArguments(len(arr), 3); result != "$3" {
		panic("at the disco")
	}
}

func TestGenerateSeparatedLine(t *testing.T) {
	if RepeatSeparated("?", ",", 5) != "?,?,?,?,?" {
		panic("at the disco")
	}
}
