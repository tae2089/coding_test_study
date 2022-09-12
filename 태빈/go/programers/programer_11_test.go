package coding

import "testing"

func TestStrToInt(t *testing.T) {
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
				s: "1234",
			},
			want: 1234,
		},
		{
			name: "case2",
			args: args{
				s: "-1234",
			},
			want: -1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrToInt(tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}