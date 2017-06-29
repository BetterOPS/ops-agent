package core

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"run-ci/pkg/log"
	"sync"
	"time"
)

type Supervisor struct {
	mux       sync.RWMutex
	processes map[string]*Process
	ctx       context.Context
}

func NewSupervisor(ctx context.Context) *Supervisor {
	return &Supervisor{
		processes: make(map[string]*Process),
		ctx:       ctx,
	}
}

func (sup *Supervisor) Start() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-sup.ctx.Done():
			return
		case <-t.C:
			now := time.Now()
			for name, p := range sup.processes {
				switch p.state {
				case STATE_WAITTING:
					sup.Run(p.Task)
				case STATE_FAILED, STATE_FINISHED:
					if p.Type() == TASK_LONG_TERM {
						sup.Run(p.Task)
					} else {
						d := viper.GetDuration(fmt.Sprintf("%s.interval", name))
						if p.updateAt.Add(d).Before(now) {
							sup.Run(p.Task)
						}
					}
				default:
				}
			}
		}
	}
}

func (sup *Supervisor) Run(t Task) {
	sup.mux.Lock()
	defer sup.mux.Unlock()

	if _, ok := sup.processes[t.Name()]; ok {
		p := sup.processes[t.Name()]
		if p.state == STATE_RUNNING || p.state == STATE_KILLING {
			log.Errorf("task %s is running or killing", t.Name())
			return
		}
		sup.exec(p)
		return
	}

	p := &Process{
		Task:     t,
		createAt: time.Now(),
		state:    STATE_WAITTING,
	}
	sup.processes[t.Name()] = p
	sup.exec(p)
	return
}

func (sup *Supervisor) exec(p *Process) {
	config := viper.Sub(p.Name())
	if err := p.Initialize(config); err != nil {
		log.Fatalf("initialize task %s error: %v", p.Name(), err)
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("running task %s error: %v", p.Name(), err)
				p.setState(STATE_FAILED)
			}
		}()

		if config != nil && config.InConfig("delay") {
			<-time.After(config.GetDuration("delay"))
		}
		p.setState(STATE_RUNNING)
		if err := p.Run(sup.ctx); err != nil {
			p.setState(STATE_FAILED)
		} else {
			p.setState(STATE_FINISHED)
		}
	}()
}

func (sup *Supervisor) Stop() {
	sup.mux.RLock()
	defer sup.mux.RUnlock()

	for _, p := range sup.processes {
		p.Kill()
	}
}
