package main

import (
	"encoding/binary"
	"os"
)

func main() {
	for u := uint64(0); u < 10; u++ {
		binary.Write(os.Stdout, binary.BigEndian, u)
		binary.Write(os.Stdout, binary.LittleEndian, u)
	}
}
