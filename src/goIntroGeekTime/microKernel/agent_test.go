package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollector(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *DemoCollector) Init(receiver EventReceiver) error {
	fmt.Println("initialize collector: ", c.name)
	c.evtReceiver = receiver
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector: ", c.name)
	for {
		select {
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collector: ", c.name)
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("timeout while stop collector: " + c.name)
	}
}

func (c *DemoCollector) Destory() error {
	fmt.Println("destory collector: " + c.name)
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollector("c1", "1")
	c2 := NewCollector("c2", "2")
	agt.RegisterCollector("c1", c1)
	agt.RegisterCollector("c2", c2)
	if err := agt.Start(); err != nil {
		fmt.Printf("start error: %v", err)
	}
	//fmt.Println(agt.Start())
}
