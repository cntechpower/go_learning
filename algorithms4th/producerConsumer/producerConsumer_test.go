package main

import (
	"testing"
	"time"
)

func Test_ProducerConsumer(t *testing.T) {
	pc := NewProducerConsumer(5, 5)
	go pc.StartProducer()
	go pc.StartConsumerCoordinator()
	go pc.FlowControl()
	time.Sleep(30 * time.Second)
	pc.GracefulShutdown()
}
