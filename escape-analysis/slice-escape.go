package main

func main() {
	xyz := [3]int{1, 2, 3}
	// main xyz does not escape
	xyzSlice := xyz[:]
	f(xyzSlice)
	h(xyzSlice)
}

func g(x int) int {
	return x
}

// f ints does not escape
func f(ints []int) {
	for _, i := range ints {
		g(i)
	}
}

// leaking param: ints to result result level=0
func h(ints []int) (result []int) {
	result = ints
	return result
}

// i ints does not escape
func i(ints []int) (result []int) {
	// []int literal escapes to heap
	result = []int{4, 5, 6}
	return result
}
