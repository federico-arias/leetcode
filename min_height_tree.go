package main

func findMinHeightTrees(n int, edges [][]int) []int {
	adj := map[int][]int{}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[0]], edge[0])
	}

	leaves := []int{}
	for val, node := range adj {
		if len(node) == 1 {
			leaves = append(leaves, val)
		}
	}

	tmp := []int{}

	for n > 2 {
		for _, leaf := range leaves {
			neigh := adj[leaf][0]
			adj[neigh] = removeLeaf(adj[neigh], leaf)
			if len(adj[neigh]) == 1 {
				tmp = append(tmp, neigh)
			}
			delete(adj, leaf)

			n--
		}
		leaves = tmp
		tmp = []int{}
	}
	res := []int{}
	for key, _ := range adj {
		res = append(res, key)
	}
	return res
}

func removeLeaf(list []int, leaf int) []int {
	res := []int{}
	for _, el := range list {
		if el != leaf {
			res = append(res, el)
		}
	}
	return res
}
