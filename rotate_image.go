package main

// 48
func rotateImage(a [][]int) [][]int {

	for i, rows := range a {
		for j, val := range rows {
			if j < i {
				continue
			}
			a[i][j] = a[j][i]
			a[j][i] = val
		}
	}

	cols := len(a[0])
	cmp := int(cols / 2)
	for i, rows := range a {
		for j, val := range rows {
			if j > cmp {
				continue
			}
			a[i][j], a[cols-j-1][i] = a[cols-j-1][i], val
		}
	}

	return a
}
