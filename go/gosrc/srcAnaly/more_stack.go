package main

import "fmt"

func moreStack(size int) {
	stackOrHeap := make([]int, size)
	fmt.Printf("lean: %d, cap: %d\n", len(stackOrHeap), cap(stackOrHeap))
}

func main() {
	for i := 1; i < 10; i++ {
		moreStack(2 << i)
	}
}
