package coding

import "testing"

func TestAddDivisor(t *testing.T) {
	type args struct {
		left  int
		right int
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
				left:  13,
				right: 17,
			},
			want: 43,
		},
		{
			name: "case2",
			args: args{
				left:  24,
				right: 27,
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDivisor(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("AddDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}