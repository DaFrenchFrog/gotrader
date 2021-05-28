package apiCaller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

//Call is...
func Call(url string, responseType reflect.Type) {
	httpResp, err := http.Get(url)
	fmt.Println(responseType.Name())
	if err != nil {
		log.Fatal(err)
	}

	defer httpResp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(httpResp.Body)
	var resp = responseType{}
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp success: %v \n", resp.Success)
}
