package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"s1", "s2", "s3"}
	slice2 := slice1[1:]

	slice2[1] = "s4"
	fmt.Printf("slice1: %v, slice2: %v\n", slice1, slice2)
	fmt.Printf("slice1: %p, slcie2: %p\n", &slice1, &slice2)
	slice2 = append(slice2, "josing")
	fmt.Printf("slice1: %p, slcie2: %p\n", &slice1, &slice2)
	fmt.Printf("slice1: %v, slice2: %v\n", slice1, slice2)
}
