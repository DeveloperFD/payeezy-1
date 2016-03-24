package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func CalculateHmac(message string, secret string) string {
	hmacSha256 := hmac.New(sha256.New, []byte(secret))
	hmacSha256.Write([]byte(message))
	hmac2Hex := hex.EncodeToString(hmacSha256.Sum(nil))
	hexArray := strings.Replace(hmac2Hex, "-", "", -1)
	return base64.StdEncoding.EncodeToString([]byte(hexArray))
}

func random(min, max int) int {
	// Create Random Number between provided Range
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func AuthPost(apiKey string, NONCE string, TimeStamp string, Merchant_Token string, jsonString string, base64Hash string) {
	// Set Endpoint URL
	url := "https://api-cert.payeezy.com/v1/transactions"

	// Prepare JSON Payload
	jsonBytes := []byte(jsonString)

	// Create HTTP Request and Set HTTP Headers
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("timestamp", TimeStamp)
	req.Header.Set("nonce", NONCE)
	req.Header.Set("token", Merchant_Token)
	req.Header.Set("apikey", apiKey)
	req.Header.Set("Authorization", base64Hash)
	req.Header.Set("Content-Type", "application/json")

	// Execute the Transaction Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Display Request Headers and Body
	fmt.Println("API Key:  " + apiKey + "\n" + "NONCE:  " + NONCE + "\n" + "Timestamp:  " + TimeStamp + "\n" + "Merchant Token:  " + Merchant_Token + "\n\n" + "HMAC Authorization: " + base64Hash + "\n\n" + "Request Body:  " + jsonString + "\n")

	// Display Transaction Response
	fmt.Println("Response Status:  " + resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body: ", string(body))
}

func main() {
	// Define API Keys, Tokens and Header Information
	apiKey := ""    //Add API Key
	apiSecret := "" //Add API Secret
	NONCE := fmt.Sprintf("%d", random(1000000001, 2147483647))
	TimeStamp := fmt.Sprintf("%d", time.Now().UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)))
	Merchant_Token := "" //Add Merchant Token

	// Define JSON Transaction Payload
	jsonString := "{ \"merchant_ref\": \"\", \"transaction_type\": \"authorize\", \"method\": \"credit_card\", \"amount\": \"10000\", \"currency_code\": \"USD\", \"credit_card\": {\"type\": \"visa\", \"cardholder_name\": \"Test Testingham\", \"card_number\": \"5301330008445732\", \"exp_date\": \"1216\", \"cvv\": \"664\" } , \"billing_address\": { \"city\": \"Hagerstown\", \"country\": \"US\", \"email\": \"\", \"phone\": { \"type\": \"P\", \"number\": \"\" }, \"street\": \"1 Western Maryland Parkway\", \"state_province\": \"AL\", \"zip_postal_code\": \"21740\" }    }"

	// Construct Hash Data, Create HMAC Hash and Display the Hash on screen
	hashData := apiKey + NONCE + TimeStamp + Merchant_Token + jsonString
	base64Hash := CalculateHmac(hashData, apiSecret)
	fmt.Println("API Secret:  " + apiSecret + "\n")

	// POST Transaction Request to Payeezy Server
	AuthPost(apiKey, NONCE, TimeStamp, Merchant_Token, jsonString, base64Hash)
}
