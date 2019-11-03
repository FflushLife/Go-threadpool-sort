package pool

import (
	"barrier"
	"fmt"
	"time"
	"unsafe"
)

type callBack func(unsafe.Pointer, uint64)

type Pool struct {
	cb callBack
	// Callback data
	cbStruct unsafe.Pointer
	br *barrier.Barrier
	cycle, started, end bool
	tCount uint64
}

func New(tCount uint64, cb callBack, cbs unsafe.Pointer) *Pool {
	var p Pool = Pool {
		cb: cb,
		cbStruct: cbs,
		br: barrier.New(tCount),
		cycle: true,
		started: false,
		end: false,
		tCount: tCount,
	}
	return &p
}

func (p *Pool) Start() {
	if !p.started {
		p.started = true
		for i := uint64(0); i < p.tCount; i++ {
			go p.wait(i);
		}
	}
	for !p.end{}
}

func (p *Pool) wait(n uint64) {
	//callback(cbStruct, n)
	// TODO:: learn why millisecond is required
	p.br.Before()
	fmt.Println("All threads are ready for job")
	time.Sleep(time.Second)
	p.cb(p.cbStruct, n)
	p.br.After()
	fmt.Println("All threads did the job")

	p.end = true
}

func (p *Pool) SetCallback(f callBack) {
	p.cb = f
}

func (p *Pool) SetCycle(c bool) {
	p.cycle = c
}
