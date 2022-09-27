package leetcode

import "testing"

func Test_checkTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name:"case1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:1,
					},
					Right: &TreeNode{
						Val:4,
					},
				},
			},
			want: true,
		},
		{
			name:"case2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:2,
					},
					Right: &TreeNode{
						Val:4,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTree(tt.args.root); got != tt.want {
				t.Errorf("checkTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
