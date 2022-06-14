package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)


func POST(url string, data interface{}) ([]byte, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	b, _ := json.Marshal(&data)
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}
