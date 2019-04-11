package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"shopping-list/shoppinglist"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	port := getPort()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/webhooks", webhookHandler).Methods("GET")
	r.HandleFunc("/webhooks", messageHandler).Methods("POST")

	log.Printf("Server up and running")

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal("Error in starting the server:", err)
	}

	// init db session
	shoppinglist.InitSession()

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index route on server!")
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":3500"
	}

	log.Printf("Using port %s", port)

	return port
}
