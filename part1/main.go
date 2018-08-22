package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// Message is struct for JSON
type Message struct {
	Body string
}

func main() {
	var message Message
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		message.Body = s.Text()
		enc := json.NewEncoder(os.Stdout)
		err := enc.Encode(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
