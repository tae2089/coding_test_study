package coding

import (
	"testing"
)

func TestIsPlusAandB(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{x: 1, y:2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusAandB(tt.args.x,tt.args.y); got != tt.want {
				t.Errorf("PlusAandB() = %v, want %v", got, tt.want)
			}
		})
	}
}