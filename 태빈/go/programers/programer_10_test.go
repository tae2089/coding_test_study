package coding

import (
	"reflect"
	"testing"
)

func TestFindMinValue(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{arr: []int{10}},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinValue(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMinValue() = %v, want %v", got, tt.want)
			}
		})
	}
}