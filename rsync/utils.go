package rsync

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func split(slice []byte, n int) [][]byte {
	if (len(slice) < n) {
		return [][]byte{slice}
	}
	length := len(slice) / n
	if (len(slice) % n > 0) {
		length++
	}
	slices := make([][]byte, length)
	for i, j := 0, 0; i < len(slices); i, j = i+1, j+n {
		slices[i] = slice[j:min(j+n, len(slice))]
	}
	return slices
}