package main

func main() {
	xyz := [3]int{1, 2, 3}
	xyzSlice := xyz[:]
	f(xyzSlice)
	h(xyzSlice)
}

func g(x int) int {
	return x
}

func f(ints []int) {
	for _, i := range ints {
		g(i)
	}
}

func h(ints []int) (result []int) {
	result = ints
	return result
}

func i(ints []int) (result []int) {
	result = []int{4, 5, 6}
	return result
}

func j() (result []int) {
	var x [3]int = [3]int{3, 2, 1}
	s := x[:0]
	s[0] = 4
	return
}

func k() (result [3]int) {
	result = [3]int{3, 2, 1}
	return
}

func l() (result *[3]int) {
	x := &[3]int{3, 2, 1}
	result = x
	return
}
