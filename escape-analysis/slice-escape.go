package main

func main() {
	xyz := [3]int{1, 2, 3}
	xyzSlice := xyz[:]
	f(xyzSlice)
}

func g(x int) int {
	return x
}

func f(ints []int) {
	for _, i := range ints {
		g(i)
	}
}
