package coding

import (
	"reflect"
	"testing"
)

func TestGCDandLCM(t *testing.T) {
	type args struct {
		n int
		m int
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
				n:3,
				m:12,
			},
			want: []int{3, 12},
		},
		{
			name:"case2",
			args: args{
				n:2,
				m:5,
			},
			want: []int{1,10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDandLCM(tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GCDandLCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
