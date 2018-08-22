package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
)

type Message struct {
	Body string
}

var listenAddress = flag.String("listen", "localhost:8000", "host:port to listen on")

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	dec := json.NewDecoder(c)
	for {
		var m Message
		if err := dec.Decode(&m); err != nil {
			log.Fatal(err)
		}
		fmt.Println(m.Body)
	}
}
