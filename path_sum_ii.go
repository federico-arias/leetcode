package main

// left result agg
// right result agg
// eval value and agg
// - top position
// - bottom position

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	pathSumDFS(root, targetSum, []int{}, &res)
	return res
}

func pathSumDFS(node *TreeNode, remainder int, partial []int, res *[][]int) {
	if node == nil {
		return
	}
	rmdr := remainder - node.Val
	p := append([]int{}, partial...)
	p = append(p, node.Val)
	if rmdr == 0 && node.Left == nil && node.Right == nil {
		*res = append(*res, p)
	}
	pathSumDFS(node.Left, rmdr, p, res)
	pathSumDFS(node.Right, rmdr, p, res)
}
