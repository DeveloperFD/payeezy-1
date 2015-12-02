package payeezy

type Header struct {
	Url           string `json:"url"`
	APIKey        string `json:"apikey"`
	Token         string `json:"token"`
	Authorization string `json:"authorization"`
	Nonce         string `json:"nonce"`
	Timestamp     string `json:"timestamp"`
}
