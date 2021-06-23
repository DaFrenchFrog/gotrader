package ftx

import (
	"fmt"

	"github.com/elRomano/gotrader/model"
)

//ListMarkets :
func (c Client) ListMarkets() (model.CoinListResponse, error) {
	marketResponse := model.CoinListResponse{}
	resp, err := c.get("markets")

	if err != nil {
		return marketResponse, fmt.Errorf("can't get markets: %w", err)
	}
	err = c.processResponse(resp, &marketResponse)
	if err != nil {
		return marketResponse, fmt.Errorf("can't process markets response: %w", err)
	}
	return marketResponse, nil
}

// GetMarketSummary :
func (c Client) GetMarketSummary(coin string) (model.MarketDataResponse, error) {
	coinDataResponse := model.MarketDataResponse{}
	resp, err := c.get("markets/" + coin)

	if err != nil {
		return coinDataResponse, fmt.Errorf("can't get coin list: %w", err)
	}
	err = c.processResponse(resp, &coinDataResponse)
	if err != nil {
		return coinDataResponse, fmt.Errorf("can't process coin list response: %w", err)
	}

	return coinDataResponse, nil
}

// GetMarketHistory :
func (c Client) GetMarketHistory(coin string, resolution int64, startTime int64, endTime int64) (model.CoinHistoryResponse, error) {
	var historicalPrices model.CoinHistoryResponse
	resp, err := c.get(fmt.Sprintf("markets/%v/candles?resolution=%v&start_time=%v&end_time=%v", coin, resolution, startTime, endTime))

	if err != nil {
		return historicalPrices, fmt.Errorf("error getting coin history, %w", err)
	}
	err = c.processResponse(resp, &historicalPrices)
	if err != nil {
		return historicalPrices, fmt.Errorf("can't process coin history response: %w", err)
	}

	return historicalPrices, nil
}
