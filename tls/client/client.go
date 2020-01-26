package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
)

var encrypt *bool = flag.Bool("encrypt", false, "Encrypt underlying connection")
var address *string = flag.String("address", ":5858", "Address to connect to.")
var port *int = flag.Int("port", 5858, "Port")
var serverName *string = flag.String("servername", "TestServer", "Name of server")

func main() {
	fmt.Println("client")
	flag.Parse()

	var err error
	var conn net.Conn

	if *encrypt {
		config := &tls.Config{
			ServerName:         *serverName,
			InsecureSkipVerify: true,
		}
		conn, err = tls.Dial("tcp", *address, config)
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = net.Dial("tcp", *address)
		if err != nil {
			panic(err)
		}
	}

	b := make([]byte, 1000)
	n, err := conn.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", n, "bytes:", string(b))

}
