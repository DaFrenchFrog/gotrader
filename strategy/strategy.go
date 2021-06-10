package strategy

import (
	"github.com/elRomano/gotrader/account"
	"github.com/elRomano/gotrader/cfmt"
	coinreader "github.com/elRomano/gotrader/coinReader"
	"github.com/elRomano/gotrader/model"
)

type strategy interface {
	init()
	apply(w *account.Wallet, ticker model.CoinHistoryDataTicker)
}

//StrategyRunner :
type StrategyRunner struct {
	wallet   account.Wallet
	strategy strategy
	market   model.CoinData
}

//New :
func New(s strategy, w account.Wallet) StrategyRunner {
	return StrategyRunner{
		wallet:   w,
		strategy: s,
	}
}

//Backtest :
func (s StrategyRunner) Backtest(coin string) {
	reader := coinreader.New()
	reader.LoadCoin(coin)
	s.market = reader.Market

	s.strategy.init()

	cfmt.Println(cfmt.Purple, "Starting backtest : ", s.market.Name+" -> testing ", len(s.market.History), " entries", cfmt.Neutral)

	for _, v := range s.market.History {
		s.strategy.apply(&s.wallet, v)
	}
	cfmt.Println(cfmt.Blue, "BACKTEST FINISHED : WINNINGS : ", s.wallet.WalletAmount, "$ // TRADES : ", s.wallet.TradeAmount)
}
