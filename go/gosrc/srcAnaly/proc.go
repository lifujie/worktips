package main

import (
	"fmt"
	"time"
)

func createG(a string, b int) {
	fmt.Printf("%s, %v, %p\n", a, &b, &a)
}

func main() {
	var josingStr string
	var josingInt int
	josingInt = 7
	josingStr = "justtostring"
	go createG(josingStr, josingInt)
	time.Sleep(1000 * time.Second)
}
