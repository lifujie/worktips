package main

import "fmt"

type PointAndScalar struct {
	First  *int32 // 1
	First2 int32  // 0
	Three  *bool  // 1
	Three2 bool   // 0
	Secon  *int64 // 1
	Secon2 int64  // 0
	Four   string // 1 0
	//Four2 *int32
}

func PrintSomething(pas *PointAndScalar, str string) {
	fmt.Printf("%v\n", *pas)
}

func PrintfIntPointerArray(pia []*int) {
	fmt.Printf("%v\n", pia)
}

type twoInt64 struct {
	aa int
	bb int
}

func PrintfIntStruct(pia *twoInt64) {
	fmt.Printf("%v\n", pia)
}

var (
	int32Num = int32(32)
	int64Num = int64(64)
	boolNum  = false
)

func intStruct() {
	bb1 := 1
	bb2 := 2
	bb := &twoInt64{bb1, bb2}
	go PrintfIntStruct(bb)
}

func intPointer() {
	bb1 := 1
	bb2 := 2
	bb := []*int{&bb1, &bb2, &bb1, &bb2, &bb1, &bb2, &bb1, &bb2, &bb1, &bb2, &bb1, &bb2}
	go PrintfIntPointerArray(bb)
}

func main() {
	intPointer()
	aa := PointAndScalar{
		First:  &int32Num,
		First2: int32Num,
		Three:  &boolNum,
		Three2: boolNum,
		Secon:  &int64Num,
		Secon2: int64Num,
		Four:   "PointAndScalar",
		//Four2: &int32Num,
	}
	str := "PointAndScalar"
	go PrintSomething(&aa, str)

}
