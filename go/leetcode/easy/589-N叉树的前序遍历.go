package main

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res = []int{}
	preorderTemp(root)
	return res
}

func preorderTemp(root *Node) {
	if root != nil {
		res = append(res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderTemp(root.Children[i])
		}
	}
}

func main() {

}
