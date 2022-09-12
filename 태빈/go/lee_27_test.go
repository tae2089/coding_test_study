package coding

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test1",
		// 	args: args{nums: []int{3,2,2,3},val: 3},
		// 	want: 2,
		// },
		// {
		// 	name: "test2",
		// 	args: args{nums: []int{0,1,2,2,3,0,4,2},val: 2},
		// 	want: 5,
		// },
		{
			name: "test3",
			args: args{nums: []int{1},val: 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
