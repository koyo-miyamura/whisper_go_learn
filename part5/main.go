package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"code.google.com/p/whispering-gophers/util"
)

type Message struct {
	Addr string
	Body string
}

var (
	peerAddr = flag.String("peer", "", "peer address")
	self     string
)

func main() {
	flag.Parse()
	l, err := util.Listen()
	go dial(*peerAddr)
	self = l.Addr().String()
	log.Println("Listening on", self)
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
			log.Println(err)
			return
		}
		fmt.Printf("%#v\n", m)
	}
}

func dial(address string) {
	var message Message
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		message.Addr = self
		message.Body = s.Text()
		enc := json.NewEncoder(c)
		err := enc.Encode(message)
		if err != nil {
			log.Fatal(err)
		}
	}
}
