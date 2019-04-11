package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MessengerInput struct {
	Entry []struct {
		Time      uint64 `json:"time,omitempty"`
		Messaging []struct {
			Sender struct {
				Id string `json:"id"`
			} `json:"sender,omitempty"`
			Recipient struct {
				Id string `json:"id"`
			} `json:"recipient,omitempty"`
			Timestamp uint64 `json:"timestamp,omitempty"`
			Message   *struct {
				Mid  string `json:"mid,omitempty"`
				Seq  uint64 `json:"seq,omitempty"`
				Text string `json:"text"`
			} `json:"message,omitempty"`
		} `json:"messaging"`
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("message received")
	defer r.Body.Close()

	input := new(MessengerInput)
	if err := json.NewDecoder(r.Body).Decode(input); err == nil {

		log.Println("got message:", input.Entry[0].Messaging[0].Message.Text)
		return
	}
}
