package ftx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
}

const URL = "https://ftx.com/api/"

func (Client) get(path string) (*http.Response, error) {
	resp, err := http.Get(URL + path)
	if err != nil {
		return nil, fmt.Errorf("can't create request %w", err)
	}
	return resp, nil
}

func (Client) processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error processing response %w", err)
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error unmarshaling response %w", err)
	}
	return nil
}
