package account

//Wallet :
type Wallet struct {
	StartingBalance float32
	Balance         float32
}

//New :
func New(amount float32) Wallet {
	return Wallet{
		Balance:         amount,
		StartingBalance: amount,
	}
}

//RegisterMarkets :
// func (w *Wallet) RegisterMarkets(markets []string, amount float32) {
// 	for _, m := range markets {
// 		w.Balance[m] = amount
// 		w.StartingBalance[m] = amount
// 	}
// }
