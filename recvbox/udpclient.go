package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "10.0.10.10:2252")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("0123456789"))
	if err != nil {
		log.Fatalln(err)
	}
}
