package tricks

import "testing"

func TestCamelCaseToLowerWithDelimiter(t *testing.T) {
	type args struct {
		s         string
		delimiter byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "underscore",
			args: args{
				s:         "TestCamelCaseMessage",
				delimiter: '_',
			},
			want: "test_camel_case_message",
		},
		{
			name: "minus",
			args: args{
				s:         "TestCamelCaseMessage",
				delimiter: '-',
			},
			want: "test-camel-case-message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCaseToLowerWithDelimiter(tt.args.s, tt.args.delimiter); got != tt.want {
				t.Errorf("CamelCaseToLowerWithDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinWithoutBlanks(t *testing.T) {
	if JoinTrimmed("-", "То", "Чаво", "Не", "Может", "Быть") != "То-Чаво-Не-Может-Быть" {
		panic("at the disco")
	}

	if JoinTrimmed("-", "То", "", "\t \tНе", "Может", "Быть  ") != "То-Не-Может-Быть" {
		panic("at the disco")
	}
}

func TestMD5Hash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test \"secret\"",
			args:    args{"secret"},
			want:    "5ebe2294ecd0e0f08eab7690d2a6ee69",
			wantErr: false,
		},
		{
			name:    "test empty",
			args:    args{},
			want:    "d41d8cd98f00b204e9800998ecf8427e",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MD5Hash(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrings(t *testing.T) {
	testSlice := []string{"aaa", "bbb", "aaa", "ccc", "ccc"}
	testSlice = UniqueStrings(testSlice)
	if len(testSlice) != 3 {
		panic("at the disco")
	}
}
