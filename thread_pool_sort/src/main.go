package main

import (
	"fmt"
	"pool"
	"psort"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	var sortInstance *psort.PSort
	var poolInstance *pool.Pool
	var tCount uint64 = 4
	runtime.GOMAXPROCS(int(tCount))

	fmt.Println("Start initializing...")
	target := []int{3, 4, 1, 11, 8, 7, 2, 6, 5, 0, 9}
	for i := 30000000; i > 12; i-- {
		target = append(target, i)
	}
	sortInstance = psort.New(target, tCount)
	// TODO:: try to create with nil
	poolInstance = pool.New(tCount, psort.TSort, unsafe.Pointer(sortInstance))

	start := time.Now()
	poolInstance.Start();
	fmt.Println("time=", time.Now().Sub(start))

	result := make([]int, 0)
	var sliceThreadSize float64 = float64(len(sortInstance.GetTarget())) / float64(tCount)
	for i := uint64(0); i < tCount; i++ {
		l := int(float64(i) * sliceThreadSize)
		r := int(float64(i + 1) * sliceThreadSize)
		result = psort.Merge(result, sortInstance.GetTarget()[l:r])
	}
	fmt.Println("final time=", time.Now().Sub(start))

	sortInstance = psort.New(target[:10], tCount)
	poolInstance.ChangeTask(unsafe.Pointer(sortInstance))
	poolInstance.Start()
	sliceThreadSize = float64(len(sortInstance.GetTarget())) / float64(tCount)
	for i := uint64(0); i < tCount; i++ {
		l := int(float64(i) * sliceThreadSize)
		r := int(float64(i + 1) * sliceThreadSize)
		result = psort.Merge(result, sortInstance.GetTarget()[l:r])
	}
	fmt.Println(result)
}
