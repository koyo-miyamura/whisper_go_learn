package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type Message struct {
	Body string
}

var (
	dialAddress   = flag.String("dial", "", "TCP connection address")
	listenAddress = flag.String("listen", "", "host:port to listen on")
)

func main() {
	flag.Parse()
	dial(*dialAddress)
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
			log.Print(err)
			return
		}
		fmt.Printf("%#v\n", m)
	}
}

func dial(address string) {
	var message Message
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.Print(err)
		return
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		message.Body = s.Text()
		enc := json.NewEncoder(c)
		err := enc.Encode(message)
		if err != nil {
			log.Print(err)
			return
		}
	}
}
