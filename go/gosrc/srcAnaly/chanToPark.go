package main

import "time"

func processChannel(ci chan int) {
	for {
		<-ci
	}
}

func putChannel(ci chan int) {
	for {
		time.Sleep(2)
		ci <- 2
	}
}

func main() {
	ci := make(chan int, 10)
	go putChannel(ci)
	go processChannel(ci)

	time.Sleep(1000 * time.Second)
}
