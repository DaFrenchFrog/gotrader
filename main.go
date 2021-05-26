package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
   "os"
   "coinReader"
)

type response struct {
   Success bool `json:"success"`
   Result []result `json:"result"`
}

type result struct{
   Name string `json:"name"`
   BaseCurrency string `json:baseCurrency`
   QuoteCurrency string `json:quoteCurrency`
   Type string `json:type`
   Underlying string `json:underlying`
   Enabled bool `json:"enabled"`
   Ask float32  `json:"ask"`
   Bid float32  `json:"bid"`
   Last float32  `json:"last"`
   PostOnly bool `json:"postOnly"`
   PriceIncrement float32 `json:"priceIncrement"`
   SizeIncrement float32 `json:"sizeIncrement"`
   Restricted bool `json:"restricted"`
}



func main() {
if(os.Args[0] == "read") {
	reader := coinReader
} else {
   httpResp, err := http.Get("https://ftx.com/api/markets")
   if err != nil {
      log.Fatal(err)
   }

   defer httpResp.Body.Close()
   bodyBytes, _ := ioutil.ReadAll(httpResp.Body)

   var resp = response{}
   err = json.Unmarshal(bodyBytes, &resp)
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("resp success: %v \n",  resp.Success)

   for _, r := range resp.Result{
      fmt.Printf("{\tname: %v\tpriceIncrement: %v}\n", r.Name, r.PriceIncrement)
   }
   }
}