package coding

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
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
			name: "test1",
			args: args{
				s: "abcabcbb",
			},
			want:3,
		},
		{
			name: "test2",
			args: args{
				s: "bbbb",
			},
			want:1,
		},
		{
			name: "test3",
			args: args{
				s: "pwwkew",
			},
			want:3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
