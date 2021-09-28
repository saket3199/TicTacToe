package Game

import (
	"strconv"

	"github.com/saket3199/TicTacToe/Board"
	"github.com/saket3199/TicTacToe/analyzer"
	"github.com/saket3199/TicTacToe/cell"
	"github.com/saket3199/TicTacToe/player"
)

type Game struct {
	Board.Board
	Players []player.Player
	row     uint8
	col     uint8
	c       string
}

func New(Size uint8) *Game {

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
func (g *Game) PutMark(position []uint8) uint8 {

	g.row = position[0]
	g.col = position[1]

	if g.Players[0].GetPlayerState() {
		g.c = g.Players[0].GetPlayerMark()

	} else if g.Players[1].GetPlayerState() {
		g.c = g.Players[1].GetPlayerMark()
	}

	// g.GetCell(row, col).SetCellMark(c)
	pos := g.row*g.GetBoard().GetSize() + g.col

	if g.row > g.GetBoard().GetSize()-1 || g.col > g.GetBoard().GetSize()-1 {
		return 1

	} else if g.GetBoard().GetCell(g.row, g.col).GetCellMark() != strconv.Itoa(int(pos)) {
		return 2
	}
	return 0

}
func (g *Game) TakeInput() uint8 {
	if g.Players[0].GetPlayerState() && !g.Players[1].GetPlayerState() {
		return 1
	} else if !g.Players[0].GetPlayerState() && g.Players[1].GetPlayerState() {
		return 2
	}
	return 0
}

func (g *Game) ResultAnalysis() uint8 {
	if analyzer.PlayerHasWon(g.GetBoard().GetAllCells()) == g.Players[0].GetPlayerMark() {
		return 1
	} else if analyzer.PlayerHasWon(g.GetBoard().GetAllCells()) == g.Players[1].GetPlayerMark() {

		return 2
	} else {
		if Board.IsBoardFull(g.GetBoard().GetAllCells()) {
			return 3
		} else {
			g.ChangeTurn()
		}
		return 0
	}
}

func (g *Game) SetMark() {
	g.GetBoard().GetCell(g.row, g.col).SetCellMark(g.c)
}

func (g *Game) ChangeTurn() {
	g.Players[0].SetPlayerState(!g.Players[0].GetPlayerState())
	g.Players[1].SetPlayerState(!g.Players[1].GetPlayerState())
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
