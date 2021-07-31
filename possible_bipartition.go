// 886 graphs, graph coloring
package main

func possibleBipartition(n int, dislikes [][]int) bool {
	adj := map[int][]int{}
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}

	for _, dislike := range dislikes {
		p0, p1 := dislike[0], dislike[1]
		adj[p0] = append(adj[p0], p1)
		adj[p1] = append(adj[p1], p0)
	}

	// 0 = unvisited; 1 and -1 = groups
	colors := []int{}

	for node, nodes := range adj {
		if colors[node] != 0 {
			continue
		}
		if !dfs(node, nodes, adj, colors, 1) {
			return false
		}
	}

	return true
}

func dfs(node int, nodes []int, adj map[int][]int, colors []int, color int) bool {
	// We can color it b/c we are sure it is 0
	colors[node] = color

	for _, n := range nodes {
		if colors[n] != 0 {
			if colors[n] != color {
				return false
			}
			continue
		}
		if !dfs(n, adj[n], adj, colors, color*-1) {
			return false
		}
	}
	return true
}
