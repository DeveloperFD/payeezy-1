package payeezy

import (
	"encoding/json"
	"log"
)

func AuthorizePayload(auth *Authorize) (payload []byte) {

	payload, err := json.Marshal(auth)
	if err != nil {
		log.Println(err)
	}
	log.Println("------")
	log.Println(string(payload))
	log.Println("------")

	return
}

// func PurchasePayload() (payload []byte) {

// }
// func CapturePayload() (payload []byte) {

// }
