package Game

import (
	"fmt"
	"log"

	"github.com/saket3199/TicTacToe/cell"
	"github.com/saket3199/TicTacToe/player"
)

// var player1, player2 player.Player

func GetUserName(players *[]player.Player) {
	var p1, p2 string
	fmt.Println("Let's Play Tic Tac Toe!")
	fmt.Println("Player 1, What's Your Name?")
	fmt.Scan(&p1)
	fmt.Println("Player 2, What's Your Name?")
	fmt.Scan(&p2)
	player1 := player.Player{
		PlayerName: p1,
		PlayerTurn: true,
	}
	player2 := player.Player{
		PlayerName: p2,
		PlayerTurn: false,
	}

	player.New(p2, false)
	*players = append(*players, player1)
	*players = append(*players, player2)
	// Game.Players = append(game.Players, player1)
	// Game.Players = append(game.Players, player2)
	// game.Players = append(game.Players, player1)
	// game.Players = append(game.Players, player2)
}
func GetBoardSize() int {
	fmt.Println("Enter Board Size")
	var size int
	fmt.Scan(&size)
	return size
}
func Play() {
	game := New(GetBoardSize())
	var players []player.Player

	DrawBoard(game.GetBoard().GetAllCells())
	GetUserName(&players)
	game.Players = append(game.Players, players...)
	fmt.Println(game.Players)
	for {
		DrawBoard(game.GetAllCells())
		WhoseTurn(game.TakeInput(), game.Players)
		for {
			i := game.PutMark(UserPosition())
			if i == 1 || i == 2 {
				BoardValidator(i)
			} else {
				break
			}
		}
		game.SetMark()
		i := game.ResultAnalysis()
		if i == 1 || i == 2 || i == 3 {
			PrintResult(i, game.Players)
			break
		}
	}
	DrawBoard(game.GetAllCells())

}

func DrawBoard(board [][]cell.Cell) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Print(board[i][j].GetCellMark())
		}
		fmt.Println()

	}
}

func WhoseTurn(i int, p []player.Player) {
	switch i {
	case 1:
		fmt.Println(p[0].GetPlayerName() + "'s Turn X")
	case 2:
		fmt.Println(p[1].GetPlayerName() + "'s Turn O")
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
	if _, err := fmt.Scan(&row); err != nil {
		log.Print("  Scan for row failed, due to ", err)
	}

	fmt.Println("Enter a column number (0, 1, or 2): ")
	if _, err := fmt.Scan(&col); err != nil {
		log.Print("  Scan for col failed, due to ", err)
	}

	positions := []int{}
	positions = append(positions, row)
	positions = append(positions, col)

	return positions
}
func PrintResult(i int, p []player.Player) {
	switch i {
	case 1:
		fmt.Println(p[0].GetPlayerName(), " has Won")
	case 2:
		fmt.Println(p[1].GetPlayerName(), "has Won")
	case 3:
		fmt.Println("Its a Tie ")
	default:
		fmt.Println("Error")
	}
}
