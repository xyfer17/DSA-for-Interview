func rotate(matrix [][]int) {

	// assign matrix length to variable n
	n := len(matrix)

	// swap the row values from top to bottom
	for i := 0; i < n; i++ {
		left := 0
		right := n - 1
		for left < right {
			tmp := matrix[left][i]
			matrix[left][i] = matrix[right][i]
			matrix[right][i] = tmp
			left++
			right--
		}
	}

	// transpose the matrix
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}
