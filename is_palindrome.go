package main

// O(2n)
func isPalindrome(s string) bool {
	rs := []rune(s)
	start := 0
	end := len(rs) - 1
	for start < end {
		if rs[start] < 'a' || rs[start] > 'z' {
			start++
		}
		if rs[end] < 'a' || rs[end] > 'z' {
			end--
		}
		if rs[start] != rs[end] {
			return false
		}
	}
	return true
}
