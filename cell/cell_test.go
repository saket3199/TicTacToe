package cell

import "testing"

func TestNewCell(t *testing.T) {
	expectedMark := NoMark
	actual := Newcell()

	if actual.Mark != expectedMark {
		t.Error("Actual is ", actual.Mark, "but excepted is : ", expectedMark)
	}
}

func TestCellMark(t *testing.T) {
	actualMark := "x"
	cell := Newcell()
	cell.SetCellMark(actualMark)
	if cell.GetCellMark() != actualMark {
		t.Error("Actual is ", actualMark, "but excepted is : ", cell.GetCellMark())
	}

}

func TestCellTaken(t *testing.T) {
	actual := false
	cell := Newcell()

	if cell.CellTaken() != actual {
		t.Error("Actual is ", actual, "but excepted is : ", cell.CellTaken())
	}
	cell.SetCellMark("X")
	if cell.CellTaken() == actual {
		t.Error("Actual is ", actual, "but excepted is : ", cell.CellTaken())
	}
}
