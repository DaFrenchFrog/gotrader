package model

//J'ai mi les structure de reponse dans un package a part, c'est pas obligatoire mais ca aller devenir relou au fur et a mesur d'en rajouter

type Response struct {
	Success bool `json:"success"`
}

//Succeed tells if the call succeed (c'est ca qui rempli le contrat de l'interface et qui permet de passer l'objet en param de l'api caller
func (r Response) Succeed() bool {
	return r.Success
}

type CoinListResponse struct {
	Response            //object composition (donc implemet l'interface ci dessus)
	Result   []CoinData `json:"result"`
}

type CoinDataResponse struct {
	Response //object composition (donc implemet l'interface ci dessus)
	Result   CoinData
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
