package coding

import "testing"

func TestTranslateString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				s: "one4seveneight",
			},
			want: 1478,
		},
		{
			name: "case2",
			args: args{
				s: "23four5six7",
			},
			want: 234567,
		},
		{
			name: "case3",
			args: args{
				s: "2three45sixseven",
			},
			want: 234567,
		},
		{
			name: "case4",
			args: args{
				s: "123",
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateString(tt.args.s); got != tt.want {
				t.Errorf("TranslateString() = %v, want %v", got, tt.want)
			}
		})
	}
}