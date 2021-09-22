package Game

import (
	"github.com/saket3199/TicTacToe/Board"
	"github.com/saket3199/TicTacToe/analyzer"
	"github.com/saket3199/TicTacToe/player"
)

var row, col int

type Game struct {
	Board [][]Board.Board
	Size  int

	// Player     string
	// TurnNumber int
	Players []player.Player
}

func (g *Game) NewGame(Size int) {
	g.Size = 3
	g.Board = make([][]Board.Board, 3)
	g.Players = make([]player.Player, 2)
	// for i := 0; i < Size; i++ {
	// 	g.Board = append(g.Board, "")
	// }
	// g.Board.SetSize(Size)

}
func (g *Game) GetSize() int {
	return g.Size
}
func (g *Game) SetSize(Size int) {
	g.Size = Size
}
func (g Game) GetBoard() Board.Board {
	return g.Board[g.Size][g.Size]
}
func (g *Game) PutMark(position []int) int {
	var c string
	c = ""
	if g.Players[0].GetPlayerState() {
		c = "X"
		g.Players[0].SetPlayerState(false)
		g.Players[1].SetPlayerState(true)

	} else if g.Players[1].GetPlayerState() {
		c = "O"
		g.Players[0].SetPlayerState(true)
		g.Players[1].SetPlayerState(false)

	}
	g.GetBoard().GetCells()[row][col].SetCellMark(c)
	row = position[0]
	col = position[1]
	if row < 0 || col < 0 || row > g.Size-1 || col > g.Size-1 {
		return 1

	} else if g.GetBoard().GetCells()[row][col] != g.GetBoard().GetCells()[row][col] {
		// GetCells()[g.row][g.col] != g.GetBoard().Cells[g.row][g.col] {
		return 2
	}
	return 0

}
func TakeInput() int {
	if player1.GetPlayerState() {
		player1.SetPlayerState(false)
		player2.SetPlayerState(true)
		return 1
	} else if player2.GetPlayerState() {
		player1.SetPlayerState(true)
		player2.SetPlayerState(false)
		return 2

	}
	return 0

}

func (g Game) ResultAnalysis() int {
	if analyzer.PlayerHasWon(g.GetBoard().GetCells()) == "X" {
		return 1
	} else if analyzer.PlayerHasWon(g.GetBoard().GetCells()) == "O" {

		return 2
	} else {
		if Board.BoardIsFull(g.GetBoard().GetCells()) {
			return 3
		} else {
			// p11 := g.Players[0].GetPlayerState()
			// p11 = !g.Players[0].GetPlayerState()
			g.Players[0].SetPlayerState(!g.Players[0].GetPlayerState())
		}
	}
	return 0
}

// func Askforplay() int {
// 	var moveInt int
// 	fmt.Println("Enter Pos to play: ")
// 	fmt.Scan(&moveInt)
// 	return moveInt
// }
// func (game *Game) Play(pos int) error {
// 	if game.Board[pos-1] == "" {
// 		game.Board[pos-1] = game.Player
// 		game.SwitchPlayers()
// 		game.TurnNumber += 1
// 		return nil
// 	}
// 	return errors.New("try another move")
// }
// func (game *Game) SwitchPlayers() {
// 	if game.Player == "O" {
// 		game.Player = "X"
// 		return
// 	}
// 	game.Player = "O"
// }
