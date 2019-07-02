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

func DeferUseParaStruct() {
	para := OneIntAndStringSt{
		One: 1,
		Two: "21",
	}
	defer OneParaStruct(para)
}

func DeferUseParaInt() {
	para := 1
	defer OneParaInt(para)
}

func DeferUseFloat64() {
	para := 1.0
	defer OneParaFloat64(para)
}

func DeferUseString() {
	para := "do something 123"
	defer OneParaString(para)
}

func main() {
	DeferUseParaInt()
	DeferUseFloat64()
	DeferUseString()
	DeferUseParaStruct()
}
