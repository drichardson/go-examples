package main

import (
	"os"
)

func main() {
	var x, y, z int
	array := []interface{}{x, y, z}
	g(x)
	h(x, y, z)
	h2(array)
}

func g(a interface{}) {
	switch v := a.(type) {
	case int:
		if v >= 0 {
			os.Stdout.WriteString("+int\n")
		} else {
			os.Stdout.WriteString("-int\n")
		}
	default:
		os.Stdout.WriteString("other\n")
	}
}

func h(as ...interface{}) {
	for _, a := range as {
		g(a)
	}
}

func h2(as []interface{}) {
	for _, a := range as {
		g(a)
	}
}
