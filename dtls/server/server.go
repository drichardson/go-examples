package main

import (
	"flag"
	"fmt"
	"github.com/pion/dtls"
	"net"
)

var encrypt = flag.Bool("encrypt", false, "Use DTLS to encrypt connection.")
var address = flag.String("address", ":5959", "Address to listen on.")

func main() {
	fmt.Println("server")

	flag.Parse()

	if *encrypt {
		serveDTLS()
	} else {
		servePlaintext()
	}
}

func serveDTLS() {
	cert, privateKey, err := dtls.GenerateSelfSigned()

	config := &dtls.Config{
		Certificate:        cert,
		PrivateKey:         privateKey,
		InsecureSkipVerify: true,
	}

	laddr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		panic(err)
	}
	listener, err := dtls.Listen("udp", laddr, config)

	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)

	for {
		fmt.Println("accept")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("read")
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Read %d bytes from %v. Data: %s\n", n, conn.RemoteAddr(), string(buffer[0:n]))

		fmt.Println("close")
		conn.Close()
	}
}

func servePlaintext() {
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
