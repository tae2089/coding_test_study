package coding

import (
	"reflect"
	"testing"
)

func TestFindIntArrayAboucDivisor(t *testing.T) {
	type args struct {
		arr     []int
		divisor int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				arr:     []int{5, 9, 7, 10},
				divisor: 5,
			},
			want: []int{5, 10},
		},
		{
			name: "case2",
			args: args{
				arr:     []int{2, 36, 1, 3},
				divisor: 1,
			},
			want: []int{1, 2, 3, 36},
		},
		{
			name: "case2",
			args: args{
				arr:     []int{3, 2, 6},
				divisor: 10,
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIntArrayAboucDivisor(tt.args.arr, tt.args.divisor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindIntArrayAboucDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}