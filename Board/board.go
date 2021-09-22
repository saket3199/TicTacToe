package Board

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/saket3199/TicTacToe/cell"
)

func ClearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

type Board struct {
	Board [][]cell.Cell
	Size  int
}

func (b Board) SetSize(Size int) {
	b.Size = Size
}
func (b Board) GetSize() int {
	return b.Size
}

func (b Board) GetBoard() [][]cell.Cell {
	return b.Board
}

func (b Board) GetCell(i, j int) cell.Cell {
	return b.Board[i][j]
}
func (b Board) GetCells() [][]cell.Cell {
	return b.Board
}
func GenerateBoard(b [][]cell.Cell) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			// cell := cell.Cell
			b[i][j] = cell.Cell{Mark: ""}
		}

	}
}

func BoardIsFull(b [][]cell.Cell) bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[i][j].GetCellMark() == "" {
				return false
			}
		}

	}
	return true

}
func PrintBoard(b [9]string) {
	ClearScreen()
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Print(v)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")

		}

	}
}
