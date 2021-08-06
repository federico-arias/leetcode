package main

func findFrequentTreeSum(root *TreeNode) []int {
	// every node is a sum
	sums := []int{}
	findFrequentTreeSumDFS(root, &sums)
	return mode(sums)
}

func mode(s []int) []int {
	nums := map[int]int{}
	for _, val := range s {
		nums[val]++
	}
	freq := map[int][]int{}
	max := 0
	for n, f := range nums {
		freq[f] = append(freq[f], n)
		if f > max {
			max = f
		}
	}
	//fmt.Printf("freq map %v, rep is %t\n", freq, repeated)
	return freq[max]
}

func findFrequentTreeSumDFS(node *TreeNode, ss *[]int) int {
	if node == nil {
		return 0
	}
	// sum = left + right + val
	left := findFrequentTreeSumDFS(node.Left, ss)
	right := findFrequentTreeSumDFS(node.Right, ss)
	sum := left + right + node.Val
	*ss = append(*ss, sum)
	return sum
}
