package main

func findAnagrams(s string, p string) []int {
	missing := map[rune]int{}
	diff := len([]rune(p))
	res := []int{}
	// string is a slice of bytes, we want a slice of chars (runes)
	str := []rune(s)

	for _, ch := range p {
		missing[ch]++
	}

	start := 0

	for end, _ := range s {

		if end-start >= len(p) {
			missing[str[start]]++
			if missing[str[start]] > 0 {
				diff++
			}
			start++
		}

		missing[str[end]]--
		if missing[str[end]] >= 0 {
			diff--
		}

		if diff == 0 {
			res = append(res, start)
		}
	}
	return res
}
