package cell

const (
	XMark  = "X"
	OMark  = "O"
	NoMark = "-"
)

type Cell struct {
	Mark string
}

func Newcell() Cell {
	return Cell{
		Mark: NoMark,
	}
}

func (c *Cell) SetCellMark(Mark string) {
	c.Mark = Mark
}
func (c *Cell) GetCellMark() string {
	return c.Mark
}

func (c *Cell) CellTaken() bool {
	return c.Mark != NoMark
}
