package pool

import (
	"barrier"
	"sync"
	"unsafe"
)

type callBack func(unsafe.Pointer, uint64)

type Pool struct {
	cb callBack
	cbStruct unsafe.Pointer // Callback data
	br *barrier.Barrier // Goroutines sync
	wg sync.WaitGroup // Master sync
	m sync.Mutex // For atomic check wautgroup current state
	started, locked bool
	tCount, wCount uint64
}

func New(tCount uint64, cb callBack, cbs unsafe.Pointer) *Pool {
	var p Pool = Pool {
		cb: cb,
		cbStruct: cbs,
		br: barrier.New(tCount),
		started: false,
		locked: false,
		tCount: tCount,
		wCount: 0,
	}
	return &p
}

// Must be guaranteed that previous task ended
// Needs lock after first use
func (p *Pool) Start() {
	if !p.started {
		p.started = true
		for i := uint64(0); i < p.tCount; i++ {
			go p.wait(i);
		}
	}
	p.wg.Add((int)(p.tCount))
	if (p.locked) {
		p.Unlock()
	}
	p.br.GiveTask()
	p.wg.Wait()
}

func (p *Pool) wait(n uint64) {
	for {
		p.br.Before()
		// Wait for new task
		p.cb(p.cbStruct, n)
		p.br.After()
		// Decrement wCount
		p.wg.Done()
	}
}

func (p *Pool) SetCallback(f callBack) {
	p.cb = f
}

func (p *Pool) ChangeTask(cbs unsafe.Pointer) {
	p.cbStruct = cbs;
}

func (p *Pool) Lock() {
	p.m.Lock()
	p.locked = true
}

func (p *Pool) Unlock() {
	p.locked = false
	p.m.Unlock()
}
