package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

// Message is struct for JSON
type Message struct {
	Body string
}

var (
	address = flag.String("address", "localhost:8080", "TCP connection address")
	message Message
)

func main() {
	flag.Parse()
	c, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		message.Body = s.Text()
		enc := json.NewEncoder(c)
		err := enc.Encode(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	//defer c.Close()
}
