package main

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

const (
	Waiting = iota
	Running
)

var WrongStateError = errors.New("can not take operation in current state")

type CollectorsError struct {
	CollectorsErrors []error
}

func (ce CollectorsError) Error() string {
	if ce.CollectorsErrors == nil {
		return ""
	}
	var errsInStringSlice []string
	for _, err := range ce.CollectorsErrors {
		errsInStringSlice = append(errsInStringSlice, err.Error())
	}
	return strings.Join(errsInStringSlice, ";")
}

type Event struct {
	Source  string
	Content string
}

type EventReceiver interface {
	OnEvent(e Event)
}

type Collector interface {
	Init(er EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destory() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuff    chan Event
	cancelFunc context.CancelFunc
	ctx        context.Context
	state      int
}

func (agt *Agent) EventProcessGoroutine() {
	var evtSeg [10]Event
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuff:
			case <-agt.ctx.Done():
				return
			}
			fmt.Println(evtSeg)
		}
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	return &Agent{
		collectors: map[string]Collector{},
		evtBuff:    make(chan Event, sizeEvtBuf),
		state:      Waiting,
	}
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent) startCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex
	for name, collector := range agt.collectors {
		go func(name string, c Collector, ctx context.Context) {
			err = collector.Start(ctx)
			mutex.Lock()
			defer mutex.Unlock()
			if err != nil {
				errs.CollectorsErrors = append(errs.CollectorsErrors, errors.New(name+" "+err.Error()))
			}
		}(name, collector, agt.ctx)
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destory(); err != nil {
			errs.CollectorsErrors = append(errs.CollectorsErrors, errors.New(name+" "+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) destoryCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		if err = collector.Destory(); err != nil {
			errs.CollectorsErrors = append(errs.CollectorsErrors, errors.New(name+" "+err.Error()))
		}
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	agt.state = Running
	agt.ctx, agt.cancelFunc = context.WithCancel(context.Background())
	go agt.EventProcessGoroutine()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancelFunc()
	return agt.stopCollectors()
}

func (agt *Agent) Destory() error {
	if agt.state != Waiting {
		return WrongStateError
	}
	return agt.destoryCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuff <- evt
}
