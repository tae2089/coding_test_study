package coding

import (
	"reflect"
	"testing"
)

func TestJumpToX(t *testing.T) {
	type args struct {
		x int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				x: 2,
				n: 5,
			}, //
			want: []int64{2, 4, 6, 8, 10},
		},
		{
			name: "case1",
			args: args{
				x: 4,
				n: 3,
			}, //
			want: []int64{4, 8, 12},
		},
		{
			name: "case1",
			args: args{
				x: -4,
				n: 2,
			}, //
			want: []int64{-4, -8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpToX(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JumpToX() = %v, want %v", got, tt.want)
			}
		})
	}
}