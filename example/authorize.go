package main

import (
	"github.com/giorgisio/payeezy_golang/payeezy"
	"log"
)

func main() {

	header := new(payeezy.Header)
	header.Url = "https://api-cert.payeezy.com/v1/transactions"
	header.APIKey = "y6pWAJNyJyjGv66IsVuWnklkKUPFbb0a"
	header.Token = "fdoa-a480ce8951daa73262734cf102641994c1e55e7cdf4c02b6"
	header.Authorization = "NDBkODhkMGQwZTEzYzkxNjc4OTc2MmM3ODFlZTA4M2MxNjcxZDY0ZDk5ZDFiOTdjZjkwNTFkNmY4NTUxNWRlYQ=="
	header.Nonce = "6770831660134717000"
	header.Timestamp = "1423683480051"

	authorize := new(payeezy.Authorize)
	authorize.MerchantRef = "Astonishing-Sale"
	authorize.TransactionType = "authorize"
	authorize.Method = "credit_card"
	authorize.Amount = "1299"
	authorize.CurrencyCode = "USD"
	authorize.CreditCard.Type = "visa"
	authorize.CreditCard.CardholderName = "John Smith"
	authorize.CreditCard.CardNumber = "4788250000028291"
	authorize.CreditCard.ExpDate = "1030"
	authorize.CreditCard.Cvv = "123"

	payload := payeezy.AuthorizePayload(authorize)

	response, err := header.PostTransaction(payload)
	if err != nil {
		log.Println(err)
	}

	log.Println(response)

	log.Println("TransactionStatus:", response.TransactionStatus)
	log.Println("ValidationStatus:", response.ValidationStatus)

	errs := len(response.Error.Messages)

	for i := 0; i < errs; i++ {
		log.Println("Code:", response.Error.Messages[i].Code)
		log.Println("Description:", response.Error.Messages[i].Description)
	}

}
