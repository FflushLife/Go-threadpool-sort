package pool

import (
	"barrier"
	"fmt"
	"time"
	"unsafe"
)

type callBack func(unsafe.Pointer, int)

type Pool struct {
	cb callBack
	// Callback data
	cbStruct unsafe.Pointer
	br *barrier.Barrier
	cycle, end bool
}

func (p *Pool) wait(n int) {
		//fmt.Println("Waiting...")
		//callback(cbStruct, n)
		// Learn why millisecond is required
	p.br.Before()
	fmt.Println("All threads are ready for job")
	for p.cycle {
		time.Sleep(time.Second)
	}
	p.br.After()
	fmt.Println("All threads did the job")

	p.end = true
}

func New(tCount int, cb callBack, cbs unsafe.Pointer) *Pool {
	var p Pool = Pool {
		cb: cb,
		cbStruct: cbs,
		br: barrier.New(tCount),
		cycle: true,
		end: false,
	}
	return &p
}

func (p *Pool) SetCallback(f callBack) {
	p.cb = f
}

func (p *Pool) SetCycle(c bool) {
	p.cycle = c
}
