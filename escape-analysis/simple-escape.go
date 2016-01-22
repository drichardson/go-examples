package main

func main() {
	f1()
}

func f1() *int {
	x := 1
	return &x
}
