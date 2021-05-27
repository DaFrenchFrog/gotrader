func Call(responseType struct, url string) {
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

}