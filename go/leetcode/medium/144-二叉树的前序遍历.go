package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var resp []int

func preorderTraversal(root *TreeNode) []int {
	resp = []int{}
	temp(root)
	return resp
}

func temp(root *TreeNode) {
	if root != nil {
		resp = append(resp, root.Val)
		temp(root.Left)
		temp(root.Right)
	}
}

func main() {

}
