package main

func firstNotRepeatingCharacter(s string) string {
	rep := map[rune]int{}

	for _, ch := range s {
		rep[ch]++
	}

	for _, ch := range s {
		if rep[ch] == 1 {
			return string(ch)
		}
	}
	return "_"
}
