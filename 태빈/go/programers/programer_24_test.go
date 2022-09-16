package coding

import "testing"

func TestStrangeString(t *testing.T) {
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
				s: "try hello world",
			},
			want: "TrY HeLlO WoRlD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrangeString(tt.args.s); got != tt.want {
				t.Errorf("StrangeString() = %v, want %v", got, tt.want)
			}
		})
	}
}