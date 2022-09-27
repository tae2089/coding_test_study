package coding

import (
	"reflect"
	"testing"
)

func TestRandomSorting(t *testing.T) {
	type args struct {
		strings []string
		n       int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				strings: []string{"sun", "bed", "car"},
				n:       1,
			},
			want: []string{"car", "bed", "sun"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomSorting(tt.args.strings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomSorting() = %v, want %v", got, tt.want)
			}
		})
	}
}
