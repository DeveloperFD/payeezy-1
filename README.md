## Payeezy Golang Library

No longer maintained (January 2016).

UPDATE (March 24, 2016) -- Refer: payeezy2/payeezy.go - use if you have to. (Very much don't recommend FirstData/Payeezy)

You're probably going to run into "code":"403", "message":"HMAC validation Failure" (FirstData/Payeezy support wasn't much help with the issue).

```
	// Test URL
	url := "https://api-cert.payeezy.com/v1/transactions"

	// Type
	contentType := "Content-Type:application/json"

	//apiKey = Available on your payeezy account
	apiKey = "y6pWAJNyJyjGv66IsVuWnklkKUPFbb0a"

	//apiSecret = Available on your payeezy account
	apiSecret = ""

	//nonce = generate a random alphanumeric value
	nonce = "6770831660134717000"

	//timestamp = UTC, EPOCH, Millseconds, within 5 minutes of firstdata server time
	//Example code (Format = 1423683480051)
	t := time.Now().UTC()
	timestamp := strconv.Itoa(int(t.UnixNano() / int64(time.Millisecond)))

	//Token = Available on your payeezy account
	token = "fdoa-a480ce8951daa73262734cf102641994c1e55e7cdf4c02b6"

	// Payload
	payload := `{
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

	c := []string{apiKey, nonce, timestamp, token, payload}
	concatedMessage = strings.Join(c, "")

	key := []byte(apiSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(concatedMessage))

	authorization := base64.StdEncoding.EncodeToString(h.Sum(nil))
```

Use the above values to post to test URL

```
	// "github.com/giorgisio/payeezy/payeezy"
	header := new(payeezy.Header)

	//Set the following
	//header.Url, header.APIKey, header.Nonce, header.Timestamp, header.Token, header.Authorization

	// POST
	response, err := header.PostTransaction(payload)

```

First Data Payeezy integration 

##Transaction Details

https://developer.payeezy.com/creditcardpayment/apis/post/transactions

##Other languages API

https://github.com/payeezy/payeezy_direct_API

##Contribute

Please fork and pull request.

##Examples

See example directory

##License

The Golang implementation of Payeezy is open source and available under the MIT license.