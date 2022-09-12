package coding

import (
	"reflect"
	"testing"
)

func TestReversIntarray(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name:"case1",
			args: args{
				n: 12345,
			}, 
			want: []int{5,4,3,2,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReversIntarray(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReversIntarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
