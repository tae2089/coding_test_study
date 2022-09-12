package coding

import "testing"

func TestReversInt(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{n: 118372},
			want: 873211,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReversInt(tt.args.n); got != tt.want {
				t.Errorf("ReversInt() = %v, want %v", got, tt.want)
			}
		})
	}
}