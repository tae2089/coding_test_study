package leetcode

import "testing"

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{Val:3},
						Right: &TreeNode{Val:7},
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{Val:18},
					},
				},
				low: 7,
				high: 15,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
