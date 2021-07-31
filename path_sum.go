package main

func pathSum(root Node, sum int) bool {
	if root == nil {
		return false
	}
	if dfs(root, sum-root.Val) {
		return true
	}
	return false

}

func dfs(node Node, diff int) {
	if diff == 0 {
		return true
	}
	if node == nil {
		return false
	}
	if dfs(root, sum-root.Val) {
		return true
	}
	return false
}
