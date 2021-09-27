package Game

import (
	"github.com/saket3199/TicTacToe/Board"
	"github.com/saket3199/TicTacToe/analyzer"
	"github.com/saket3199/TicTacToe/cell"
	"github.com/saket3199/TicTacToe/player"
)

type Game struct {
	Board.Board
	Players []player.Player
	row     int
	col     int
	c       string
}

func New(Size int) *Game {

	game := Game{
		*Board.New(Size),
		make([]player.Player, 0, 2),
		0,
		0,
		cell.NoMark,
	}
	Board.GenerateBoard(*game.GetBoard().GetBoard())
	return &game
}

func (g *Game) GetBoard() *Board.Board {
	return &g.Board
}
func (g *Game) PutMark(position []int) int {

	g.row = position[0]
	g.col = position[1]

	if g.Players[0].GetPlayerState() {
		g.c = cell.XMark

	} else if g.Players[1].GetPlayerState() {
		g.c = cell.OMark
	}

	// g.GetCell(row, col).SetCellMark(c)

	if g.row < 0 || g.col < 0 || g.row > g.GetBoard().GetSize()-1 || g.col > g.GetBoard().GetSize()-1 {
		return 1

	} else if g.GetBoard().GetCell(g.row, g.col).GetCellMark() != cell.NoMark {
		return 2
	}
	return 0

}
func (g *Game) TakeInput() int {
	if g.Players[0].GetPlayerState() && !g.Players[1].GetPlayerState() {
		return 1
	} else if !g.Players[0].GetPlayerState() && g.Players[1].GetPlayerState() {
		return 2
	}
	return 0
}

func (g *Game) ResultAnalysis() int {
	if analyzer.PlayerHasWon(g.GetBoard().GetAllCells()) == cell.XMark {
		return 1
	} else if analyzer.PlayerHasWon(g.GetBoard().GetAllCells()) == cell.OMark {

		return 2
	} else {
		if Board.IsBoardFull(g.GetBoard().GetAllCells()) {
			return 3
		} else {
			g.Players[0].SetPlayerState(!g.Players[0].GetPlayerState())
			g.Players[1].SetPlayerState(!g.Players[1].GetPlayerState())
		}
	}
	return 0
}

func (g *Game) SetMark() {
	g.GetBoard().GetCell(g.row, g.col).SetCellMark(g.c)
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
