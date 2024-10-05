
func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			val := board[i][j]
			if val == '.' {
				continue
			}

			//row check
			for r := 0; r < 9; r++ {
				if r != i && board[r][j] == val {

					return false

				}

			}

			//column check
			for c := 0; c < 9; c++ {
				if c != j && board[i][c] == val {

					return false

				}

			}

			// 3x3 box check

			row := i / 3
			col := j / 3

			for k := row * 3; k < row*3+3; k++ {
				for m := col * 3; m < (col*3)+3; m++ {

					if k != i && m != j && board[k][m] == val {
						return false
					}

				}
			}

		}
	}

	return true
}