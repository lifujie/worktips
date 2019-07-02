package main

import "fmt"

const (
	bitPointer    = 1
	bitPointerAll = 15
	bitScan       = 16
	bitScanAll    = 240
	heapBitsShift = 1
)

func main() {
	fmt.Printf("%d\n", (bitPointer | bitScan | bitPointer<<heapBitsShift))
}
