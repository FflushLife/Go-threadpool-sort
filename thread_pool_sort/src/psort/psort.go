package psort

import (
	"fmt"
	"unsafe"
)

type PSort struct {
	target []int
	tCount int
}

// Per thread sort
func TSort(uInstance unsafe.Pointer, n int) {
	fmt.Printf("Callback %d\n", n)

	var psInstance PSort = *(*PSort)(uInstance)
	var sliceThreadSize float64 = float64(len(psInstance.target)) / float64(psInstance.tCount)
	// TODO:: redo
	fmt.Println(psInstance.target[int(float64(n) * sliceThreadSize):int(float64(n + 1) * sliceThreadSize)])

	fmt.Println(psInstance.tCount - n, sliceThreadSize)
}

func New(target []int, tCount int) PSort {
	fmt.Println("Sort inited")
	return PSort{target, tCount}
}
