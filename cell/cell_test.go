package cell

import "testing"

func TestNewCell(t *testing.T) {
	expectedMark := NoMark
	actual := Newcell()

	if actual.Mark != expectedMark {
		t.Error("Actual is ", actual.Mark, "but excepted is : ", expectedMark)
	}
}
