package main

import (
	"github.com/giorgisio/payeezy_golang/payeezy"
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
	authorize.CreditCard.CardType = "visa"
	authorize.CreditCard.CardHolderName = "John Smith"
	authorize.CreditCard.CardNumber = "4788250000028291"
	authorize.CreditCard.CardExpiry = "1030"
	authorize.CreditCard.CardCvv = "123"

	payload := payeezy.AuthorizePayload(authorize)

	header.PostTransaction(payload)

}
