package coding

import (
	"reflect"
	"testing"
)

func TestLotto(t *testing.T) {
	type args struct {
		lottos   []int
		win_nums []int
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
				lottos: []int{44, 1, 0, 0, 31, 25},
				win_nums: []int{31, 10, 45, 1, 6, 19},
			},
			want: []int{3,5},
		},
		{
			name:"case2",
			args: args{
				lottos: []int{0, 0, 0, 0, 0, 0},
				win_nums: []int{38, 19, 20, 40, 15, 25},
			},
			want: []int{1,6},
		},
		{
			name:"case3",
			args: args{
				lottos: []int{45, 4, 35, 20, 3, 9},
				win_nums: []int{20, 9, 3, 45, 4, 35},
			},
			want: []int{1,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lotto(tt.args.lottos, tt.args.win_nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lotto() = %v, want %v", got, tt.want)
			}
		})
	}
}
