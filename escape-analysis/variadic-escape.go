package main

func main() {
	var x, y, z int
	f(x, y, z)
}

func g(x int) int {
	return x
}

func f(ints ...int) {
	for _, i := range ints {
		g(i)
	}
}
