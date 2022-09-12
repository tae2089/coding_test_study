package coding

import "testing"

func TestIsSqrt(t *testing.T) {
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
			args: args{
				n:121,
			},
			want:144,
		},
		{
			name: "case2",
			args: args{
				n:3,
			},
			want:-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSqrt(tt.args.n); got != tt.want {
				t.Errorf("IsSqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
