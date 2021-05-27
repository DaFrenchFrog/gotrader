package coinReader

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CoinReader struct {
	url string
}

func New(url string) CoinReader {
	return CoinReader{
		url: url,
	}
}

func (c CoinReader) Url() string {
	return c.url
}

func (c *CoinReader) SetUrl(url string) {
	c.url = url
}

func (c CoinReader) Read() {
	httpResp, err := http.Get(c.url)
	if err != nil {
		log.Fatal(err)
	}

	defer httpResp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(httpResp.Body)
	bodyString := string(bodyBytes)
	fmt.Printf(bodyString)
	//    var resp = response{}
	//    err = json.Unmarshal(bodyBytes, &resp)
	//    if err != nil {
	//       log.Fatal(err)
	//    }
	//    fmt.Printf("resp success: %v \n",  resp.Success)

	//    for _, r := range resp.Result{
	//       fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", r.Name, r.PriceIncrement)
	//    }
}
