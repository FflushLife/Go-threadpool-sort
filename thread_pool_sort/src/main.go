package main

import (
	"fmt"
	"pool"
	"psort"
	"time"
	"unsafe"
)

func main() {
	var sortInstance psort.PSort
	var tCount int = 5

	fmt.Println("Start initializing...")
	sortInstance = psort.New([]int{3, 4, 1, 8, 7}, tCount)
	pool.SetCallback(psort.TSort)
	go pool.Create(tCount, unsafe.Pointer(&sortInstance))

	time.Sleep(time.Second * 3)
	pool.SetCycle(false)
	time.Sleep(time.Second)
}
