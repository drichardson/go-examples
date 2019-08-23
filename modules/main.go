package main

import (
	"fmt"
	"github.com/drichardson/hello"
	helloV2 "github.com/drichardson/hello/v2"
	helloV3 "github.com/drichardson/hello/v3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(helloV2.HelloV2())
	fmt.Println(helloV2.Echo("hi"))
	fmt.Println(helloV3.EchoV3("hi"))
}
