package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
)

var encrypt *bool = flag.Bool("encrypt", false, "Encrypt underlying connection")
var port *int = flag.Int("port", 5858, "Port")

func main() {
	fmt.Println("client")
	flag.Parse()

	var err error
	var conn net.Conn

	if *encrypt {
		config := &tls.Config{}
		conn, err = tls.Dial("tcp", fmt.Sprintf(":%d", *port), config)
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = net.Dial("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			panic(err)
		}
	}

	b := make([]byte, 1000)
	n, err := conn.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", n, "bytes:", b)

}
