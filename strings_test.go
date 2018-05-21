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
