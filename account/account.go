package account

//Wallet :
type Wallet struct {
	Balance map[string]float32
}

//New :
func New() Wallet {
	return Wallet{
		Balance: map[string]float32{},
	}
}

//RegisterMarkets :
func (w *Wallet) RegisterMarkets(markets []string, amount float32) {
	for _, m := range markets {
		w.Balance[m] = amount
	}
}
