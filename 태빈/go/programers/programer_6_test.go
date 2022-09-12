package coding

import "testing"

func TestIsHashard(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				n: 10,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				n: 123,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHashard(tt.args.n); got != tt.want {
				t.Errorf("IsHashard() = %v, want %v", got, tt.want)
			}
		})
	}
}