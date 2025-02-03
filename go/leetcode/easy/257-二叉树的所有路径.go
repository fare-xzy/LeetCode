package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var resString []string

func binaryTreePaths(root *TreeNode) []string {
	resString = make([]string, 0)
	result := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{result}
	}
	if root.Left != nil {
		recursion(root.Left, result)
	}
	if root.Right != nil {
		recursion(root.Right, result)
	}
	return resString
}

func recursion(root *TreeNode, result string) {
	current := strconv.Itoa(root.Val)
	if root.Left != nil {
		recursion(root.Left, result+"->"+current)
	}
	if root.Right != nil {
		recursion(root.Right, result+"->"+current)
	}
	if root.Left == nil && root.Right == nil {
		resString = append(resString, result+"->"+current)
	}
}
