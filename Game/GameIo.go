package Game

import (
	"fmt"

	"github.com/saket3199/TicTacToe/cell"
	"github.com/saket3199/TicTacToe/player"
)

var player1, player2 player.Player
var p1, p2 string

func GetUserName() {
	fmt.Println("Let's Play Tic Tac Toe!")
	fmt.Println("Player 1, What's Your Name?")
	fmt.Scanln(&p1)
	fmt.Println("Player 2, What's Your Name?")
	fmt.Scanln(&p2)
	player1 = player.Player{
		PlayerName: p1,
		PlayerTurn: true,
	}
	player2 = player.Player{
		PlayerName: p2,
		PlayerTurn: false,
	}

}
func Play() {

	DrawBoard(new(Game).GetBoard().GetCells())
	GetUserName()
	for {
		DrawBoard(new(Game).GetBoard().GetCells())
		//			int j = game.takeInput();
		//			if (j == 1 || j == 2) {
		WhoseTurn(TakeInput())
		//			}
		for {
			i := new(Game).PutMark(UserPosition())
			if i == 1 || i == 2 {
				BoardValidator(i)
			} else {
				break
			}
		}
		// Game = new(Game)
		i := new(Game).ResultAnalysis()
		if i == 1 || i == 2 || i == 3 {
			PrintResult(i)
			break
		}
	}
	DrawBoard(new(Game).GetBoard().GetCells())

}

func DrawBoard(board [][]cell.Cell) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Println(board[i][j].GetCellMark())

		}
		fmt.Println()

	}
}
func WhoseTurn(i int) {
	switch i {
	case 1:
		fmt.Println(player1.GetPlayerName() + "'s Turn X")
	case 2:
		fmt.Println(player2.GetPlayerName() + "'s Turn X")
	}
}

func BoardValidator(i int) {
	switch i {
	case 1:
		fmt.Println("This Position is Out of Bounds. Try Again !")
	case 2:
		fmt.Println("Someone has Already Made a Move. Try Again ! ")
	default:
		fmt.Println("Error")
	}
}
func UserPosition() []int {
	var row, col int
	fmt.Println("Enter a row number (0, 1, or 2): ")
	fmt.Scan(&row)
	fmt.Println("Enter a column number (0, 1, or 2): ")
	fmt.Scan(&col)
	var positions []int
	positions = append(positions, row)
	positions = append(positions, col)
	return positions
}
func PrintResult(i int) {
	switch i {
	case 1:
		fmt.Println(player1.GetPlayerName(), " has Won")
	case 2:
		fmt.Println(player2.GetPlayerName(), "has Won")
	case 3:
		fmt.Println("Its a Tie ")
	default:
		fmt.Println("Error")
	}
}
