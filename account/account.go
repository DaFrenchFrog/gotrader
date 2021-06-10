package account

//Wallet :
type Wallet struct {
	WalletAmount float32
	Trading      bool
	TradeAmount  int
	Opening      float32
}

//New :
func New(amount float32) Wallet {
	return Wallet{
		WalletAmount: amount,
	}
}
