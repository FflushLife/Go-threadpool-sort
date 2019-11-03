package main

import (
	"fmt"
	"pool"
	"psort"
	//"time"
	"unsafe"
)

func main() {
	var sortInstance *psort.PSort
	var poolInstance *pool.Pool
	var tCount uint64 = 5

	fmt.Println("Start initializing...")
	sortInstance = psort.New([]int{3, 4, 1, 8, 7, 2, 6, 5, 0, 9}, tCount)
	// TODO:: try to create with nil
	poolInstance = pool.New(tCount, psort.TSort, unsafe.Pointer(sortInstance))
	poolInstance.Start();
}
