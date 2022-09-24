package leetcode


func rangeSumBST(root *TreeNode, low int, high int) int {
    arr := make([]int,0)
	var total int
	isNodeInRange(root, low, high,&arr)
	for _, value := range arr {
		total += value
	}
	return total
}

func isNodeInRange(node *TreeNode, low, high int, arr *[]int){
	if node == nil {
		return
	}
	isNodeInRange(node.Left,low,high,arr)
	if node.Val >=low && node.Val <=high {
		*arr = append(*arr, node.Val)
	}
	isNodeInRange(node.Right,low,high,arr)
}