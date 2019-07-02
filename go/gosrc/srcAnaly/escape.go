package main

import "fmt"

func main() {
	var largeHeap [1024 * 1024 * 16]byte
	hw := "hello world"
	_ = largeHeap
	fmt.Printf("%s", hw)
}
