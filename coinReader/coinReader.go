package coinReader

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
)

func coinReader()  {
httpResp, err := http.Get("https://ftx.com/api/markets/")
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