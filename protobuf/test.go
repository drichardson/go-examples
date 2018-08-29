package main

import (
	"flag"
	pb "github.com/drichardson/go-examples/protobuf/tutorial"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

var outfile = flag.String("out", "", "output filename")

func main() {
	flag.Parse()

	if len(*outfile) == 0 {
		log.Fatalln("out not set.")
	}

	book := &pb.AddressBook{}
	book.People = []*pb.Person{
		{Email: "me@example.com", Id: 123, Name: "Doug Richardson"},
		{Email: "you@example.com", Id: 123, Name: "Home Slice"},
	}

	out, err := proto.Marshal(book)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile(*outfile, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

}
