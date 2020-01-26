package main

import (
	"flag"
	"fmt"
	//"github.com/pion/dtls"
	"net"
)

var encrypt *bool = flag.Bool("encrypt", false, "Use DTLS to encrypt connection.")
var address *string = flag.String("address", ":5959", "Address to listen on.")

func main() {
	fmt.Println("server")

	flag.Parse()

	pc, err := net.ListenPacket("udp", *address)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)

	for {
		n, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Read %d bytes from %v. Data: %s\n", n, addr, string(buffer))
	}
}
