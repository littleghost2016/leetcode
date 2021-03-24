package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return nil
	}
	p := []*TreeNode{root}
	for i := 0; len(p) > 0; i++ {
		result = append(result, []int{})
		q := []*TreeNode{}
		for j := 0; j < len(p); j++ {
			result[i] = append(result[i], p[j].Val)
			if p[j].Left != nil {
				q = append(q, p[j].Left)
			}
			if p[j].Right != nil {
				q = append(q, p[j].Right)
			}
		}

		p = q
	}

	return result
}

// 一遍过
