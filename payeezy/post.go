package payeezy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Header) PostTransaction(payload []byte) (response *Response, err error) {

	req, err := http.NewRequest("POST", h.Url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", h.APIKey)
	req.Header.Set("token", h.Token)
	req.Header.Set("Authorization", h.Authorization)
	req.Header.Set("nonce", h.Nonce)
	req.Header.Set("timestamp", h.Timestamp)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	response = new(Response)
	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	return
}
