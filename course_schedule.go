package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	req := map[int][]int{}
	for _, r := range prerequisites {
		req[r[0]] = r[1]
	}
	visited := map[int]int{}

	for req0, rs := range req {

	}

}

func canFinishDFS(req *map[int]int, visited map[int]int) {
}
