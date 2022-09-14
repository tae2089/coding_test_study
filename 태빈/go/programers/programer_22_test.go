package coding

import (
	"reflect"
	"testing"
)

func TestAddArray(t *testing.T) {
	type args struct {
		arr1 [][]int
		arr2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name:"case1",
			args: args{
				arr1: [][]int{{1,2},{2,3}},
				arr2: [][]int{{3,4},{5,6}},
			},
			want: [][]int{{4,6},{7,9}},
		},
		{
			name:"case2",
			args: args{
				arr1: [][]int{{1},{2}},
				arr2: [][]int{{3},{4}},
			},
			want: [][]int{{4},{6}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
