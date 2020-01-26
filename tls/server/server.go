package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
)

var encrypt *bool = flag.Bool("encrypt", false, "Encrypt underlying connection")
var port *int = flag.Int("port", 5858, "Port")

func main() {
	fmt.Println("Server")

	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	if *encrypt {
		config := &tls.Config{}
		listener = tls.NewListener(listener, config)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection.", err)
			continue
		}
		defer conn.Close()

		n, err := conn.Write([]byte("Hello from Server"))
		if err != nil {
			log.Println("Error writing bytes to connection.", err)
			continue
		}

		fmt.Println("Wrote", n, "bytes")
	}
}
