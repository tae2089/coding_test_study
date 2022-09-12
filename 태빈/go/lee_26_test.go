package coding

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name :"test1",
			args : args{nums: []int{1,1,2}}, 
			want : 2,
		},
		{
			name :"test2",
			args : args{nums: []int{0,0,1,1,1,2,2,3,3,4}}, 
			want : 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
