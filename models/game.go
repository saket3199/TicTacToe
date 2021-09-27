package models

import (
	"github.com/saket3199/TicTacToe/Board"
	"github.com/saket3199/TicTacToe/player"
)

type Game struct {
	Board.Board
	Players []player.Player
	row     int
	col     int
	c       string
}
