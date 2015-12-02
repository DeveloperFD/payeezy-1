package main

import (
	"encoding/json"
	"github.com/giorgisio/payeezy/payeezy"
	"log"
)

func main() {

	var err error

	// HEADER
	header := new(payeezy.Header)
	headJson := `
	{
		"url": "https://api-cert.payeezy.com/v1/transactions",
		"apikey": "y6pWAJNyJyjGv66IsVuWnklkKUPFbb0a",
		"token": "fdoa-a480ce8951daa73262734cf102641994c1e55e7cdf4c02b6",
		"authorization": "NDBkODhkMGQwZTEzYzkxNjc4OTc2MmM3ODFlZTA4M2MxNjcxZDY0ZDk5ZDFiOTdjZjkwNTFkNmY4NTUxNWRlYQ==",
		"nonce": "6770831660134717000",
		"timestamp": "1423683480051"
	}
	`
	json.Unmarshal([]byte(headJson), header)
	if err != nil {
		log.Println(err)
	}

	//AUTHORIZE Data
	authorize := new(payeezy.Authorize)
	authJson := `
	{
		"merchant_ref": "Astonishing-Sale",
		"transaction_type": "authorize",
		"method": "credit_card",
		"amount": "1299",
		"currency_code": "USD",
		"credit_card": {
			"type": "visa",
			"cardholder_name": "John Smith",
			"card_number": "4005519200000004",
			"exp_date": "1020",
			"cvv": "123"
		}
	}`
	err = json.Unmarshal([]byte(authJson), authorize)
	if err != nil {
		log.Println(err)
	}

	// Construct payload
	payload := payeezy.AuthorizePayload(authorize)

	//Post Transaction
	response, err := header.PostTransaction(payload)
	if err != nil {
		log.Println(err)
	}

	errs := len(response.Error.Messages)
	for i := 0; i < errs; i++ {
		log.Println("Code:", response.Error.Messages[i].Code)
		log.Println("Description:", response.Error.Messages[i].Description)
	}

	log.Println(response)
	log.Println("TransactionStatus:", response.TransactionStatus)
	log.Println("ValidationStatus:", response.ValidationStatus)

}
