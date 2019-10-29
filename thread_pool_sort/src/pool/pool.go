package pool

import (
	"barrier"
	"fmt"
	"time"
	"unsafe"
)

type cb func(unsafe.Pointer, int)
var cycle, end = true, false
var callback cb
var cbStruct unsafe.Pointer
var br *barrier.Barrier

func wait(n int) {
		//fmt.Println("Waiting...")
		//callback(cbStruct, n)
		// Learn why millisecond is required
		br.Before()
		fmt.Println("All threads are ready for job")
	for cycle {
		time.Sleep(time.Second)
	}
	br.After()
	fmt.Println("All threads did the job")

	end = true
}

func Create(tCount int, cbs unsafe.Pointer) {
	cbStruct = cbs
	br = barrier.New(tCount)
	for i := 0; i < tCount; i++ {
		go wait(i)
	}
	fmt.Printf("%d pools have been created\n", tCount)
	for !end {}
	fmt.Println("End")
}

func SetCallback(f cb) {
	callback = f
}

func SetCycle(c bool) {
	cycle = c
}
