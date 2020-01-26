package main

import (
	"fmt"
	//"github.com/pion/dtls"
	"flag"
	"net"
)

var encrypt *bool = flag.Bool("encrypt", false, "Use DTLS to encrypt connection.")
var address *string = flag.String("address", ":5959", "Address to listen on.")

func main() {
	fmt.Println("client")

	flag.Parse()

	raddr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	n, err := conn.Write([]byte(`Hello World UDP`))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Write %d bytes.\n", n)
}
