package strategy

import "time"

//Position :
type Position struct {
	PositionType string
	open         bool
	OpeningPrice float32
	ClosingPrice float32
	Quantity     float32
	DateStart    time.Time
	DateEnd      time.Time
	Winnings     float32
}

//NewPosition :
func NewPosition(t string, opening float32, quantity float32, date time.Time) Position {
	return Position{
		open:         true,
		PositionType: t,
		OpeningPrice: opening,
		Quantity:     quantity,
		DateStart:    date,
	}
}

//Close :
func (p *Position) Close(closingPrice float32, date time.Time) (float32, bool) {
	p.DateEnd = date
	p.ClosingPrice = closingPrice
	p.open = false
	if p.PositionType == "long" {
		p.Winnings = (p.Quantity * p.ClosingPrice) - (p.Quantity * p.OpeningPrice)
		// return p.Quantity * closingPrice, closingPrice > p.OpeningPrice
		return p.Quantity * p.OpeningPrice, closingPrice > p.OpeningPrice
	} else {
		p.Winnings = (p.Quantity * p.OpeningPrice) - (p.Quantity * p.ClosingPrice)
		return (p.Quantity * p.OpeningPrice) + (closingPrice-p.OpeningPrice)*p.Quantity, closingPrice < p.OpeningPrice
	}

}
