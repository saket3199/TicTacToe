package Board

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/saket3199/TicTacToe/cell"
)

func ClearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

type Board struct {
	Board [][]cell.Cell
	Size  uint8
}

func (b *Board) SetSize(Size uint8) {
	b.Size = Size
}
func (b *Board) GetSize() uint8 {
	return b.Size
}

var board Board

func New(size uint8) *Board {
	board.SetSize(size)
	slice := make([][]cell.Cell, 0, size)
	var i uint8
	for i = 0; i < size; i++ {
		row := make([]cell.Cell, size)
		for j := range row {
			fmt.Scan(&row[j])
		}
		slice = append(slice, row)

	}
	board = Board{
		slice,
		board.Size,
	}

	fmt.Println(slice)
	return &board
}

func (b *Board) GetBoard() *[][]cell.Cell {
	return &b.Board
}

func (b *Board) GetCell(i, j uint8) *cell.Cell {
	return &b.Board[i][j]
}
func (b Board) GetAllCells() [][]cell.Cell {
	return b.Board
}
func GenerateBoard(b [][]cell.Cell) {
	count := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			// cell := cell.Cell
			b[i][j].SetCellMark(strconv.Itoa(count))
			count++
		}

	}
}

func IsBoardFull(b [][]cell.Cell) bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[i][j].GetCellMark() == cell.NoMark {
				return false
			}
		}

	}
	return true

}

//board size validation done
// no same player name done
// ask player mark done
// first 5 dosen't send to anayalzer

// func PrintBoard(b [9]string) {
// 	ClearScreen()
// 	for i, v := range b {
// 		if v == "" {
// 			fmt.Printf(" ")
// 		} else {
// 			fmt.Print(v)
// 		}

// 		if i > 0 && (i+1)%3 == 0 {
// 			fmt.Printf("\n")
// 		} else {
// 			fmt.Printf("|")

// 		}

// 	}
// }
