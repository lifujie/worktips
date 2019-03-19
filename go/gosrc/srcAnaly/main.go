package main

import "fmt"

const (
	heapArenaBytes       = 67108864 // 64M
	logHeapArenaBytes    = 26
	heapArenaBitmapBytes = 2097152 // 4KB
	arenaL1Bits          = 0
	arenaL2Bits          = 22
	arenaBits            = 22
	arenaBaseOffset      = 140737488355328 // 128G
)

type heapArena struct {
	bitmap [heapArenaBitmapBytes]byte
}

type arenaIdx uint

func (i arenaIdx) l1() uint {
	if arenaL1Bits == 0 {
		// Let the compiler optimize this away if there's no
		// L1 map.
		return 0
	} else {
		return uint(i) >> arenaL1Shift
	}
}

func (i arenaIdx) l2() uint {
	if arenaL1Bits == 0 {
		return uint(i)
	} else {
		return uint(i) & (1<<arenaL2Bits - 1)
	}
}

func main() {
	var arenas [1 << arenaL1Bits]*[1 << arenaL2Bits]*heapArena
	var a = 23
	for {
		fmt.Printf("%d\n", a)
	}
}
