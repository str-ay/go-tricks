# Go-Tricks

## CamelCaseToLowerWithDelimiter
CamelCaseToLowerWithDelimiter convert ThisText to this-text or this_text

## JoinTrimmed
JoinTrimmed makes Join with trim all blank spaces like [ ], \t, "\n"

## MD5Hash
MD5Hash generates MD5 string

## RepeatSeparated
RepeatSeparated repeats string with delimiter
RepeatSeparated("?", ",", 4) -> "?,?,?,?"

## UniqueStrings
UniqueStrings returns unique strings

## QueryArguments
QueryArguments is useful for dynamic SQL:
in(" + QueryArguments(len(ids)) + ") -> in($1,$2,$3,$4,$5)

## QuestionMarks
QuestionMarks is useful for dynamic SQL:
in(" + QuestionMarks(len(ids)) + ") -> in(?,?,?,?,?)
