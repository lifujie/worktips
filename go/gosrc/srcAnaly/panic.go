package main

import "fmt"

func throwPanic() {
	panic("throwPanic")
}

func throwParaPanic() {
	panic("throwParaPanic")
}

func main() {
	defer func(para string) {
	}("para")

	defer func(para string) {
		if err := recover(); err != "" {
			fmt.Printf("recovery failure %s\n", err)
		}
	}("para")
	throwPanic()
	throwParaPanic()
}
