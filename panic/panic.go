package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	e := errOnPanic()
	fmt.Println("2", e)

}

func errOnPanic() (result error) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case error:
				fmt.Println("Recovered an error object", t)
				result = t
			case string:
				fmt.Println("Recovered a string object", t)
				result = fmt.Errorf("%v", t)
			default:
				fmt.Println("Recovered an unknown object", t)
				result = fmt.Errorf("panicked with unknown type")
			}
		}
	}()

	fmt.Println("a")
	willPanic()
	fmt.Println("b")
	return nil
}

func willPanic() {
	panic("I panicked!")
	//panic(fmt.Errorf("an error panic"))
	//panic(1)
}
