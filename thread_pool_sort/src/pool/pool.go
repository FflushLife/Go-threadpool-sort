package pool

import (
	"barrier"
	//"fmt"
	"sync"
	"time"
	"unsafe"
)

type callBack func(unsafe.Pointer, uint64)

type Pool struct {
	cb callBack
	// Callback data
	cbStruct unsafe.Pointer
	br *barrier.Barrier
	wg sync.WaitGroup
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
			p.wg.Add(1)
			go p.wait(i);
		}
	}
	//p.wg.Wait()
	for !p.br.WorkIsDone(){time.Sleep(time.Millisecond)}
}

func (p *Pool) wait(n uint64) {
	// TODO:: learn why millisecond is required
	p.br.Before()
	p.cb(p.cbStruct, n)
	p.br.After()
	p.wg.Done()

	// TODO:: fix the error: p.end is set by the only goroutine
	//p.end = true
}

func (p *Pool) SetCallback(f callBack) {
	p.cb = f
}

func (p *Pool) SetCycle(c bool) {
	p.cycle = c
}
