package pool

import (
	"fmt"
	"time"
	"unsafe"
)

type cb func(unsafe.Pointer, int)
var cycle, end = true, false
var callback cb
var cbStruct unsafe.Pointer

func wait(n int) {
	for cycle {
		fmt.Println("Waiting...")
		callback(cbStruct, n)
		time.Sleep(time.Second)
	}
	end = true
}

func Create(tCount int, cbs unsafe.Pointer) {
	cbStruct = cbs
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
