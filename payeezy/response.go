package payeezy

/*
	Response from authorize requests
*/
type Response struct {
	CorrelationID     string `json:"correlation_id"`
	TransactionStatus string `json:"transaction_status"`
	ValidationStatus  string `json:"validation_status"`
	TransactionType   string `json:"transaction_type"`
	TransactionID     string `json:"transaction_id"`
	TransactionTag    string `json:"transaction_tag"`
	Method            string `json:"method"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Cvv2              string `json:"cvv2"`
	Token             struct {
		TokenType string `json:"token_type"`
		TokenData struct {
			Value string `json:"value"`
		} `json:"token_data"`
	} `json:"token"`
	Card struct {
		Type           string `json:"type"`
		CardholderName string `json:"cardholder_name"`
		CardNumber     string `json:"card_number"`
		ExpDate        string `json:"exp_date"`
	} `json:"card"`
	BankRespCode    string `json:"bank_resp_code"`
	BankMessage     string `json:"bank_message"`
	GatewayRespCode string `json:"gateway_resp_code"`
	GatewayMessage  string `json:"gateway_message"`
	Error           struct {
		Messages []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"messages"`
	} `json:"Error"`
}
