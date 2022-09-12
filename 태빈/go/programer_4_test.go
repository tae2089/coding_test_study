package coding

import (
	"reflect"
	"testing"
)

func TestSumDivision(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 12,
			},
			want: 28,
		},
		{
			name: "case2",
			args: args{
				n: 5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumDivision(tt.args.n); got != tt.want {
				t.Errorf("SumDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}

