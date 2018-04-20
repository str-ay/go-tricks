package tricks

import (
	"testing"
)

func TestJoinWithoutBlanks(t *testing.T) {
	if JoinTrimmed("-", "То", "Чаво", "Не", "Может", "Быть") != "То-Чаво-Не-Может-Быть" {
		panic("at the disco")
	}

	if JoinTrimmed("-", "То", "", "\t \tНе", "Может", "Быть  ") != "То-Не-Может-Быть" {
		panic("at the disco")
	}
}

func TestGenerateQuestionMarks(t *testing.T) {
	if QuestionMarks(10) != "?,?,?,?,?,?,?,?,?,?" {
		panic("at the disco")
	}
}

func TestGenerateSeparatedLine(t *testing.T) {
	if RepeatSeparated("?", ",", 5) != "?,?,?,?,?" {
		panic("at the disco")
	}
}
