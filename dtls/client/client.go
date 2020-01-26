package main

import (
	"flag"
	"fmt"
	"github.com/pion/dtls"
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

	var conn net.Conn

	if *encrypt {

		conn, err = dtls.Dial("udp", raddr, &dtls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = net.DialUDP("udp", nil, raddr)
		if err != nil {
			panic(err)
		}
	}

	n, err := conn.Write([]byte(`Hello World UDP`))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Write %d bytes.\n", n)

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}
