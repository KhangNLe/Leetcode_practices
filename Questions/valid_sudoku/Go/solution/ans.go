package solution

func IsValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]struct{})
	cols := make(map[int]map[byte]struct{})
	boxes := make(map[int]map[byte]struct{})

	for i := range len(board) {
		for j := range len(board) {
			char := board[i][j]

			if char == '.' {
				continue
			}

			if res, ok := rows[i]; ok {
				if _, ok = res[char]; ok {
					return false
				}
			} else {
				rows[i] = make(map[byte]struct{})
			}

			if res, ok := cols[j]; ok {
				if _, ok = res[char]; ok {
					return false
				}
			} else {
				cols[j] = make(map[byte]struct{})
			}

			boxPos := (i/3)*3 + (j / 3)

			if res, ok := boxes[boxPos]; ok {
				if _, ok = res[char]; ok {
					return false
				}
			} else {
				boxes[boxPos] = make(map[byte]struct{})
			}

			rows[i][char] = struct{}{}
			cols[j][char] = struct{}{}
			boxes[boxPos][char] = struct{}{}
		}
	}
	return true
}
