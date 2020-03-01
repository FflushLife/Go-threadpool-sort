package barrier

import "sync"

type Barrier struct {
	c uint64
	n uint64
	m sync.Mutex
	before chan uint64
	after chan uint64
	taskGiven chan bool
}

func New(n uint64) *Barrier {
	b := Barrier {
		n: n,
		before:	make(chan uint64, n),
		after: make(chan uint64, n),
		taskGiven: make(chan bool),
	}
	return &b
}

func (b *Barrier) Before() {
	b.m.Lock()
	b.c += 1
	if b.c == b.n {
		<-b.taskGiven
		for i := uint64(0); i < b.n; i++ {
			b.before <- 1
		}
	}
	b.m.Unlock()
	<-b.before
}

func (b *Barrier) After() {
	b.m.Lock()
	b.c -= 1
	if b.c == 0 {
		for i := uint64(0); i < b.n; i++ {
			b.after <- 1
		}
	}
	b.m.Unlock()
	<-b.after
}

func (b* Barrier) GiveTask() {
	b.taskGiven<-true
}
