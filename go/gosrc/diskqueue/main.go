package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

func sLogFunc(lvl LogLevel, f string, args ...interface{}) {
	fmt.Printf(f+"\n", args...)
}

func main() {
	dq := New("test", ".", 1024, 1, 10, 1, 10*time.Second, sLogFunc)
	go func() {
		msg := "base"
		for i := 0; i < 10; i++ {
			//fmt.Printf("2234\n")
			dq.Put([]byte(msg + strconv.Itoa(i)))
		}
	}()

	//time.Sleep(10 * time.Second)

	go func() {
		msg := "base"
		for i := 10; i < 20; i++ {
			//fmt.Printf("2233\n")
			dq.Put([]byte(msg + strconv.Itoa(i)))
		}
	}()

	go func() {

		for {
			select {
			case msg := <-dq.ReadChan():
				fmt.Printf("Read Msg: %s\n", msg)
			}
		}
	}()

	go func() {

		for {
			select {
			case msg := <-dq.ReadChan():
				fmt.Printf("Read Msg: %s\n", msg)
			}
		}
	}()

	fmt.Printf("be: %d\n", binary.BigEndian)
	time.Sleep(10 * time.Minute)
}
