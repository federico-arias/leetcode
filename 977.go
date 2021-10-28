package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func sortedSquares(nums []int) []int {
	l := len(nums)
	if l == 0 || l == 1 {
		return nums
	}
	solution = make([]int, len(nums))
	nPtr := 0
	pPtr := 0

	for i, v := range nums {
		if v >= 0 {
			pPtr = i
			nPtr = i - 1
			break
		}
	}

	for pPtr < l && nPtr >= 0 {
		sq0 := nums[pPtr] * nums[pPtr]
		sq1 := nums[nPtr] * nums[nPtr]
		if sq0 > sq1 {
			solution = append(solution, sq1)
			pPtr++
			continue
		}
		solution = append(solution, sq0)
		nPtr--
	}

	for nPtr >= 0 {
		sq1 := nums[nPtr] * nums[nPtr]
		solution = append(solution, sq1)
		nPtr--
	}

	for pPtr < l {
		sq0 := nums[pPtr] * nums[pPtr]
		solution = append(solution, sq0)
		nPtr--
	}

	return solution
}
