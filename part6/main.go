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

	self = l.Addr().String()
	log.Println("Listening on", self)

	go dial(*peerAddr)
	go readInput()

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

var peer = make(chan Message)

func readInput() {
	var message Message
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		message.Addr = self
		message.Body = s.Text()
		peer <- message
	}
	if err := s.Err(); err != nil {
		log.Println(err)
		return
	}
}

func dial(addr string) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(addr, err)
		return
	}
	defer c.Close()

	e := json.NewEncoder(c)
	for m := range peer {
		err := e.Encode(m)
		if err != nil {
			log.Println(addr, err)
			return
		}
	}
}
