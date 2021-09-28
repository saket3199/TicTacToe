package Board

import (
	"testing"

	"github.com/saket3199/TicTacToe/cell"
)

func TestNew(t *testing.T) {
	b := New(3)
	for i := 0; i < len(b.Board); i++ {
		for j := 0; j < len(b.Board); j++ {
			if b.Board[i][j].GetCellMark() != "" {
				t.Error("Actual mark is ", b.Board[i][j].GetCellMark(), "but excepted is : ", cell.NoMark)
			}
		}

	}
}

func TestGenerateBoard(t *testing.T) {
	b := New(3)
	GenerateBoard(b.Board)
	for i := 0; i < len(b.Board); i++ {
		for j := 0; j < len(b.Board); j++ {
			if b.Board[i][j].GetCellMark() != cell.NoMark {
				t.Error("Actual mark is ", b.Board[i][j].GetCellMark(), "but excepted is : ", cell.NoMark)
			}
		}

	}

}

func TestIsBoardFull(t *testing.T) {
	b := New(3)
	GenerateBoard(b.Board)
	value := IsBoardFull(b.Board)
	if value != false {
		t.Error("Actual Value is ", value, "but excepted is : ", false)

	}
}
