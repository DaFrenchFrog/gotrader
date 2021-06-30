package strategy

//Position :
type Position struct {
	PositionType string
	open         bool
	OpeningPrice float32
	Quantity     float32
}

//NewPosition :
func NewPosition(t string, opening float32, quantity float32) Position {
	return Position{
		PositionType: t,
		OpeningPrice: opening,
		Quantity:     quantity,
	}
}

//Close :
func (p *Position) Close(closingPrice float32) (float32, bool) {
	return p.Quantity * closingPrice, closingPrice > p.OpeningPrice
}
