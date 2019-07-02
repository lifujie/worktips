package main

import "fmt"

func OneParaInt(para int) {
	fmt.Printf("Print a int %d\n", para)
}

func OneParaFloat64(para float64) {
	fmt.Printf("Print a int %f\n", para)
}

func OneParaString(para string) {
	fmt.Printf("Print a int %s\n", para)
}

type OneIntAndStringSt struct {
	One int
	Two string
}

func OneParaStruct(para OneIntAndStringSt) {
	fmt.Printf("Print a struct %v\n", para)
}

func CreateGUseParaStruct() {
	para := OneIntAndStringSt{
		One: 1,
		Two: "21",
	}
	go OneParaStruct(para)
}

func CreateGUseParaInt() {
	para := 1
	go OneParaInt(para)
}

func CreateGUseFloat64() {
	para := 1.0
	go OneParaFloat64(para)
}

func CreateGUseString() {
	para := "do something 123"
	go OneParaString(para)
}

func main() {
	CreateGUseParaInt()
	CreateGUseFloat64()
	CreateGUseString()
	CreateGUseParaStruct()
}
