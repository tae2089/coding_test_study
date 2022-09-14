package coding

import "testing"

func TestSortSentance(t *testing.T) {
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
				s: "Zbcdefg",
			},
			want: "gfedcbZ",
		},
		{
			name: "case2",
			args: args{
				s: "Za",
			},
			want: "aZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSentance(tt.args.s); got != tt.want {
				t.Errorf("SortSentance() = %v, want %v", got, tt.want)
			}
		})
	}
}