package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

func mul(numbers []int) int {
	var v int
	for _, n := range numbers {
		v *= n
	}
	return v
}

const (
	windowSize = 200000
	msgCount   = 1000000
)

type (
	message []byte
	buffer  [windowSize]message
)

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(b *buffer, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func domsg() {
	var b buffer
	for i := 0; i < msgCount; i++ {
		pushMsg(&b, i)
	}
	fmt.Println("Worst push time: ", worst)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	numbers := []int{1, 2, 3, 3, 4, 5, 5, 6, 6}
	go add(numbers)
	go mul(numbers)
	go domsg()
	time.Sleep(100 * time.Second)
}
