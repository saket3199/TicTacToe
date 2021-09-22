package analyzer

import "github.com/saket3199/TicTacToe/cell"

func PlayerHasWon(board [][]cell.Cell) string {
	for i := 0; i < len(board); i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0].GetCellMark() != "" {
			return board[i][0].GetCellMark()
		}
	}

	for j := 0; j < len(board); j++ {
		if board[0][j] == board[1][j] && board[1][j] == board[2][j] && board[0][j].GetCellMark() != "" {
			return board[0][j].GetCellMark()
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0].GetCellMark() != "" {
		return board[0][0].GetCellMark()
	}
	if board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[2][0].GetCellMark() != "" {
		return board[2][0].GetCellMark()
	}

	return ""
}

// func CheckForWinner(b [9]string, n int) (bool, string) {

// 	test := false
// 	i := 0

// 	//horizantel test
// 	for i < 9 {
// 		test = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
// 		if !test {
// 			i += 3
// 		} else {
// 			return true, b[i]
// 		}
// 	}
// 	i = 0
// 	//vertical test
// 	for i < 3 {
// 		test = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
// 		if !test {
// 			i += 1
// 		} else {
// 			return true, b[i]
// 		}
// 	}
// 	//diagonal 1 test
// 	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
// 		return true, b[i]
// 	}
// 	//diagonal 2 test
// 	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
// 		return true, b[i]
// 	}
// 	if n == 9 {
// 		return true, ""
// 	}
// 	return false, ""
// }
