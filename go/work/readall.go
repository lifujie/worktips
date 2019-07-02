package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

func sizeOf() {
	var aa int64 = 1231804808015813550
	var bb = int(334241)
	fmt.Printf("%d, %d\n", unsafe.Sizeof(aa), unsafe.Sizeof(bb))
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func main() {
	sizeOf()
	read3("/Users/l00277880/selfcode/go/goproj/lean/lfj/worktips/go/work/ca.key")
}
