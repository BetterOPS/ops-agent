package core

import (
	"ops-agent/lib/log"
	"sync"
	"time"
)

type Process struct {
	sync.Mutex
	Task
	state    int
	createAt time.Time
	updateAt time.Time
}

func (p *Process) setState(state int) {
	if state < STATE_WAITTING || state > STATE_KILLING {
		log.Fatalf("%v", INVALID_STATE)
	}
	p.Lock()
	defer p.Unlock()
	p.state = state
	p.updateAt = time.Now()
}
