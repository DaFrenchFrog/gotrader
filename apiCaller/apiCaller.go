package apiCaller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Call(url string, responseType struct) {
	httpResp, err := http.Get(baseURL)
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
	fmt.Printf("resp success: %v \n", resp.Success)
	return ???
}
