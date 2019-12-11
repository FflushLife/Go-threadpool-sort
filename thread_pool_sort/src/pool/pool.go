package pool

import (
	"barrier"
	"sync"
	"unsafe"
)

type callBack func(unsafe.Pointer, uint64)

type Pool struct {
	cb callBack
	// Callback data
	cbStruct unsafe.Pointer
	br *barrier.Barrier
	wg sync.WaitGroup
	cycle, started, end, taskChanged bool
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
		taskChanged: true,
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
	p.wg.Wait()
	p.br.ResetWork()
}

func (p *Pool) wait(n uint64) {
	for {
		p.br.Before()
		// Wait for new task
		if (!p.taskChanged) {
			p.br.After()
		} else {
			p.cb(p.cbStruct, n)
			p.br.After()
			p.taskChanged = false
			p.wg.Done()
		}
	}
}

func (p *Pool) SetCallback(f callBack) {
	p.cb = f
}

func (p *Pool) SetCycle(c bool) {
	p.cycle = c
}

func (p *Pool) ChangeTask(cbs unsafe.Pointer) {
	p.cbStruct = cbs;
}
