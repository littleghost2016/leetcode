package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorder(root, &result)
	return result
}

// 这里为什么要传*[]int而不是[]int。
// 因为preorder函数里面有append操作，有对[]int的len和cap函数操作的地方，所以要传地址
// 如果只有值的修改和读取（即不对len和cap进行修改则可以改用直接传[]int
func preorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorder(root.Left, result)
	preorder(root.Right, result)
}
