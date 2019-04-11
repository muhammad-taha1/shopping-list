package main

import (
	"log"
	"strings"
)

func parseAndHandleMessage(message string) {
	message = strings.TrimSpace(message)
	if message == "" {
		return
	}

	// if message contains ':', then anything before ':' is the category name.
	// otherwise items go to default category which is Miscellaneous
	category := "Miscellaneous"
	if strings.Contains(message, ":") {
		category = strings.Split(message, ":")[0]
		message = message[strings.IndexByte(message, ':')+1:]
	}

	log.Printf("Adding new item(s) with category: " + category)

	items := strings.Split(message, "\n")
}
