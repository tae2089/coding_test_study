package coding

import "testing"

func TestFindCenterChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				s: "abcde",
			},
			want: "c",
		},
		{
			name: "case2",
			args: args{
				s: "qwer",
			},
			want: "we",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCenterChar(tt.args.s); got != tt.want {
				t.Errorf("FindCenterChar() = %v, want %v", got, tt.want)
			}
		})
	}
}