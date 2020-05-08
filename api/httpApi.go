package api

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// GetFromAPI 调用HTTP API
// method: GET,POST
// apiURL: HTTP address
// bs: parameters
// userAgent: HTTP HEAD
// 6028b9923a303cb582c27acd05baae15
func GetFromAPI(method string, apiURL string, bs []byte, userAgent string) (interface{}, error) {
	var req io.Reader
	if len(bs) > 0 {
		req = bytes.NewBuffer([]byte(bs))
	}
	client := &http.Client{}
	request, err := http.NewRequest(method, apiURL, req)
	if err != nil {
		return nil, err
	}
	if len(userAgent) > 0 {
		request.Header.Add("User-Agent", userAgent)
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	//request.Header.Add("charset", "utf-8")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get get method
func Get(apiURL string) (interface{}, error) {
	return GetFromAPI("GET", apiURL, []byte{}, "")
}

// Post post method
func Post(apiURL string, bs []byte) (interface{}, error) {
	return GetFromAPI("POST", apiURL, bs, "")
}
