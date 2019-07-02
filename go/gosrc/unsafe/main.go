package main

import (
	"fmt"
	"unsafe"
)

// StrsSt 字符串结构体
type StrsSt struct {
	Data *int8
	Len  int64
}

// StringToStruct 字符串转为结构体
func StringToStruct(para string) {
	ss := (*StrsSt)(unsafe.Pointer(&para))
	fmt.Printf("0x%x,%d\n", ss.Data, ss.Len)
}

func main() {
	para := "aaaaa"
	StringToStruct(para)
}
