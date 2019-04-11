package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

var tk Config

// var v = getToken()

// var _ = json.Unmarshal([]byte(v), &tk)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	tk.readYml()
	token := tk.VerifyToken
	tokenTrue := r.URL.Query().Get("hub.verify_token")
	hubChallenge := r.URL.Query().Get("hub.challenge")

	log.Printf(base64.StdEncoding.EncodeToString([]byte(token)))

	if tokenTrue == token {
		log.Printf("Tokens matched!")
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, hubChallenge)
	} else {
		log.Printf("Token didnt match!")
		fmt.Fprintf(w, "Tokens don't match!")
	}
}
