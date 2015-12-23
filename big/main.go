package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, ok := big.NewInt(0).SetString("19585729239412394192341892346159815151987591298356123958123565899", 10)
	if !ok {
		fmt.Errorf("Failed to set string")
	}
	fmt.Println("a is", a)
	r := new(big.Int)
	b := big.NewInt(999999)
	r.Mul(a, b)
	fmt.Println("r is", r)
	r.Div(r, b)
	fmt.Println("r is", r)
	r.Mul(r, b)

	//a.Mul(a, b)
	fmt.Println("GCD of a=", a, "and b=", b)
	r.GCD(nil, nil, a, b)
	fmt.Println("GCD is ", r)

}
