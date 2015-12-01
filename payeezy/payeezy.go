package payeezy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Header struct {
	Url           string
	APIKey        string
	Token         string
	Authorization string
	Nonce         string
	Timestamp     string
}

type Authorize struct {
	MerchantRef     string `json:"merchant_ref"`
	TransactionType string `json:"transaction_type"`
	Method          string `json:"method"`
	Amount          string `json:"amount"`
	CurrencyCode    string `json:"currency_code"`
	CreditCard      struct {
		CardType       string `json:"type"`
		CardHolderName string `json:"cardholder_name"`
		CardNumber     string `json:"card_number"`
		CardExpiry     string `json:"exp_date"`
		CardCvv        string `json:"cvv"`
	} `json:"credit_card"`
}

type Response struct {
}

func (h *Header) PostTransaction(payload []byte) {

	req, err := http.NewRequest("POST", h.Url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", h.APIKey)
	req.Header.Set("token", h.Token)
	req.Header.Set("Authorization", h.Authorization)
	req.Header.Set("nonce", h.Nonce)
	req.Header.Set("timestamp", h.Timestamp)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Status)
	log.Println(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	return
}

func AuthorizePayload(auth *Authorize) (payload []byte) {

	payload, err := json.Marshal(auth)
	if err != nil {
		log.Println(err)
	}

	return
}

// func PurchasePayload() (payload []byte) {

// }
// func CapturePayload() (payload []byte) {

// }
