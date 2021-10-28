package main

import (
	"fmt"
	"strconv"
)

func sumNumbers(root *TreeNode) int {
	ns := [][]int{}
	sumNumbersDFS(root, []int{}, &ns)
	return sum(ns)
}

func sum(ns [][]int) int {
	s := 0
	for _, n := range ns {
		str := ""
		for _, i := range n {
			str = str + strconv.Itoa(i)
		}
		i, _ := strconv.Atoi(str)
		s = s + i
	}
	fmt.Printf("slice %v, ret %d\n", ns, s)
	return s
}

func sumNumbersDFS(root *TreeNode, partial []int, ns *[][]int) {
	if root == nil {
		return
	}
	partial = append(partial, root.Val)
	if root.Left == nil && root.Right == nil {
		n := make([]int, len(partial))
		copy(n, partial)
		*ns = append(*ns, n)
	}
	//fmt.Printf("node %v, ret %d\n", root.Val, s)
	sumNumbersDFS(root.Left, partial, ns)
	sumNumbersDFS(root.Right, partial, ns)
}
