package model

type Response struct {
	Success bool `json:"success"`
}

//Succeed tells if the call succeed (c'est ca qui rempli le contrat de l'interface et qui permet de passer l'objet en param de l'api caller
func (r Response) Succeed() bool {
	return r.Success
}

type CoinListResponse struct {
	Response
	Result []CoinData `json:"result"`
}

type CoinDataResponse struct {
	Response
	Result CoinData
}

type CoinHistoryResponse struct {
	Response
	Result []CoinHistoryData
}
type CoinHistoryData struct {
	Close     float32 `json:"close"`
	High      float32 `json:"high"`
	Low       float32 `json:"low"`
	Open      float32 `json:"open"`
	ClockTime float32 `json:"time"`
	StartTime string  `json:"startTime"`
	Volume    float32 `json:"volume"`
}

type CoinData struct {
	Name           string  `json:"name"`
	BaseCurrency   string  `json:"baseCurrency"`
	QuoteCurrency  string  `json:"quoteCurrency"`
	Type           string  `json:"type"`
	Underlying     string  `json:"underlying"`
	Enabled        bool    `json:"enabled"`
	Ask            float32 `json:"ask"`
	Bid            float32 `json:"bid"`
	Last           float32 `json:"last"`
	PostOnly       bool    `json:"postOnly"`
	PriceIncrement float32 `json:"priceIncrement"`
	SizeIncrement  float32 `json:"sizeIncrement"`
	Restricted     bool    `json:"restricted"`
}

// Color get color
func Color(c string) string {

	switch c {
	case "red":
		return string("\033[31m")
	case "green":
		return string("\033[32m")
	case "yellow":
		return string("\033[33m")
	case "blue":
		return string("\033[34m")
	case "purple":
		return string("\033[35m")
	case "cyan":
		return string("\033[36m")
	case "white":
		return string("\033[37m")
	default:
		return string("\033[0m")
	}
}
