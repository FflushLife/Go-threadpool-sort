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
	var tCount int = 5

	fmt.Println("Start initializing...")
	sortInstance = psort.New([]int{3, 4, 1, 8, 7, 2, 6, 5, 0, 9}, tCount)
	poolInstance = pool.New(tCount, psort.TSort, unsafe.Pointer(&sortInstance))
	fmt.Println(poolInstance)
}
