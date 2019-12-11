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
	m sync.Mutex
	started, taskChanged, Locked bool
	tCount uint64
}

func New(tCount uint64, cb callBack, cbs unsafe.Pointer) *Pool {
	var p Pool = Pool {
		cb: cb,
		cbStruct: cbs,
		br: barrier.New(tCount),
		started: false,
		Locked: false,
		taskChanged: true,
		tCount: tCount,
	}
	return &p
}

// Must be guaranteed that previous task ended
func (p *Pool) Start() {
	if !p.started {
		p.started = true
		for i := uint64(0); i < p.tCount; i++ {
			p.wg.Add(1)
			go p.wait(i);
		}
	} else {
		p.wg.Add((int)(p.tCount))
	}
	if (p.Locked) {
		p.m.Unlock()
	}
	p.wg.Wait()
}

func (p *Pool) wait(n uint64) {
	for {
		p.br.Before()
		// Wait for new task
		p.m.Lock()
		if (!p.taskChanged) {
			p.m.Unlock()
			p.br.After()
		} else {
			p.m.Unlock()
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

func (p *Pool) ChangeTask(cbs unsafe.Pointer) {
	p.cbStruct = cbs;
	p.taskChanged = true
}

func (p *Pool) Lock() {
	p.m.Lock()
}

func (p *Pool) Unlock() {
	p.m.Lock()
}
