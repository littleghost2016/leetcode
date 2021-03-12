package main

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	treeString := "3,9,20,1,1,15,7,1"
	rootNode := createTree(treeString)
	result := maxDepth(rootNode)
	correctResult := 4
	if result != correctResult {
		t.Error(result, "结果错误")
	}
}

// 根据字符串生成对应的树
func createTree(treeString string) *TreeNode {
	length := len(treeString)
	parsedNumber := 0
	treeValueSlice := []int{}
	for i := 0; i < length; {
		// 是数字
		if '0' <= treeString[i] && treeString[i] <= '9' {
			parsedNumber = parsedNumber*10 + int(treeString[i]-'0')
			if i == length-1 {
				treeValueSlice = append(treeValueSlice, parsedNumber)
			}
			i++
			//不是数字
		} else {
			// 逗号，逗号则添加上一个数字
			if treeString[i] == ',' {
				treeValueSlice = append(treeValueSlice, parsedNumber)
				parsedNumber = 0
				i++
				// 是null的第一个字母n，
			} else if treeString[i] == 'n' {
				parsedNumber = -1
				i += 4
			}
		}
	}

	treeNodeSlice := make([]*TreeNode, 0)
	rootNode := &TreeNode{}
	// treeNodeSlice = append(treeNodeSlice, rootNode)
	for i := 0; i < len(treeValueSlice); i++ {
		if treeValueSlice[i] != -1 {
			tempTreeNode := &TreeNode{}
			tempTreeNode.Val = treeValueSlice[i]
			if len(treeNodeSlice) == 0 {
				treeNodeSlice = append(treeNodeSlice, tempTreeNode)
				rootNode = tempTreeNode
			} else if treeNodeSlice[0].Left == nil {
				treeNodeSlice[0].Left = tempTreeNode
				treeNodeSlice = append(treeNodeSlice, tempTreeNode)
			} else if treeNodeSlice[0].Right == nil {
				treeNodeSlice[0].Right = tempTreeNode
				treeNodeSlice = append(treeNodeSlice, tempTreeNode)
				treeNodeSlice = treeNodeSlice[1:]
			}
		}
	}
	// fmt.Println(treeValueSlice)

	// printTree(rootNode)

	return rootNode
}

func printTree(root *TreeNode) {
	fmt.Println(root.Val)
	// fmt.Println(root.Val, &root, root.Left, root.Right)
	if root.Left != nil {
		printTree(root.Left)
	}
	if root.Right != nil {
		printTree(root.Right)
	}

}
