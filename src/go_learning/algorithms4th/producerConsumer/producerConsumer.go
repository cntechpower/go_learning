package main

import (
	"fmt"
	"sync"
	"time"
)

const MaxPendingMessageCount = 10
const MinPendingMessageCount = 5
const MaxProducerSleepIntervalInMillisecond = 32768
const MinProducerSleepIntervalInMillisecond = 128

// ProducerConsumer
type ProducerConsumer struct {
	closingMu                    sync.RWMutex
	closing                      bool
	producerStopped              chan struct{}
	consumerStopped              chan struct{}
	currentPendingMessageCountMu sync.RWMutex
	currentPendingMessageCount   int

	//used producerSleepInterval to adaptive
	producerSleepIntervalMu            sync.RWMutex
	producerSleepIntervalInMillisecond int
	//used maxConsumerGoroutineCount for concurrency control
	maxConsumerGoroutineCount chan struct{}
	messageQueue              chan time.Time
}

func NewProducerConsumer(maxConsumerGoroutineCount, maxMessageBuffer int) *ProducerConsumer {
	pc := &ProducerConsumer{
		closingMu:                          sync.RWMutex{},
		closing:                            false,
		producerStopped:                    make(chan struct{}, 1),
		consumerStopped:                    make(chan struct{}, 1),
		producerSleepIntervalMu:            sync.RWMutex{},
		producerSleepIntervalInMillisecond: MinProducerSleepIntervalInMillisecond,
		currentPendingMessageCountMu:       sync.RWMutex{},
		currentPendingMessageCount:         0,
		maxConsumerGoroutineCount:          make(chan struct{}, maxConsumerGoroutineCount),
		messageQueue:                       make(chan time.Time, maxMessageBuffer),
	}
	return pc
}

//StartProducer start produce message.
//It support flow control by `Binary Exponential Backoff`
func (pc *ProducerConsumer) StartProducer() {
	for {
		if pc.closing {
			//graceful shutdown.
			//do not produce new message.
			fmt.Println("producer closed")
			pc.producerStopped <- struct{}{}
			return
		}

		go pc.addMessageToQueue()
		pc.producerSleepIntervalMu.RLock()
		sleepInterval := pc.producerSleepIntervalInMillisecond
		pc.producerSleepIntervalMu.RUnlock()
		time.Sleep(time.Duration(sleepInterval) * time.Millisecond)
	}
}

func (pc *ProducerConsumer) addMessageToQueue() {
	pc.currentPendingMessageCountMu.Lock()
	pc.currentPendingMessageCount++
	pc.currentPendingMessageCountMu.Unlock()
	pc.messageQueue <- time.Now()
}

func (pc *ProducerConsumer) getPendingMessageCount() int {
	pc.currentPendingMessageCountMu.Lock()
	defer pc.currentPendingMessageCountMu.Unlock()
	return pc.currentPendingMessageCount
}

func (pc *ProducerConsumer) StartConsumerCoordinator() {
	for {
		if pc.closing {
			//graceful shutdown.
			//make sure messageQueue is empty.
			pendingCount := pc.getPendingMessageCount()
			if len(pc.messageQueue) == 0 && pendingCount == 0 {
				//no further message, close messageQueue.
				close(pc.messageQueue)
				//make sure current working goroutine finish.
				for len(pc.maxConsumerGoroutineCount) != 0 {
					fmt.Printf("consumer coordinator closed, still need waiting %v working goroutine finish\n", len(pc.maxConsumerGoroutineCount))
					time.Sleep(2 * time.Second)
				}
				close(pc.maxConsumerGoroutineCount)
				pc.consumerStopped <- struct{}{}
				return
			} else {
				fmt.Printf("consumer coordinator is closing, but need waiting %v message to start process\n", len(pc.messageQueue)+pendingCount)
			}
		}
		// if maxConsumerGoroutineCount is not full then start new goroutine.
		pc.maxConsumerGoroutineCount <- struct{}{}
		go pc.consumerProcess()
	}

}

func (pc *ProducerConsumer) consumerProcess() {
	defer func() {
		<-pc.maxConsumerGoroutineCount
	}()
	t, ok := <-pc.messageQueue
	if !ok {
		//messageQueue closed, quit this goroutine.
		return
	}
	pc.currentPendingMessageCountMu.Lock()
	pc.currentPendingMessageCount--
	pc.currentPendingMessageCountMu.Unlock()
	time.Sleep(5 * time.Second)
	fmt.Printf("message send time %v, process finish time %v\n", t.Format(time.RFC3339), time.Now().Format(time.RFC3339))
	//process finish, remove current goroutine from maxConsumerGoroutineCount

}

func (pc *ProducerConsumer) FlowControl() {
	for {
		{
			//producer flow control
			//Binary Exponential Backoff
			pendingMsgCount := pc.getPendingMessageCount()
			if pendingMsgCount > MaxPendingMessageCount && pc.producerSleepIntervalInMillisecond < MaxProducerSleepIntervalInMillisecond {

				pc.producerSleepIntervalMu.Lock()
				fmt.Printf("changing producerSleepIntervalInMillisecond from (%v) to (%v*2)\n",
					pc.producerSleepIntervalInMillisecond, pc.producerSleepIntervalInMillisecond)
				pc.producerSleepIntervalInMillisecond *= 2
				pc.producerSleepIntervalMu.Unlock()
			}
			if pendingMsgCount < MinPendingMessageCount && pc.producerSleepIntervalInMillisecond > MinProducerSleepIntervalInMillisecond {
				pc.producerSleepIntervalMu.Lock()
				fmt.Printf("changing producerSleepIntervalInMillisecond from (%v) to (%v/2)\n",
					pc.producerSleepIntervalInMillisecond, pc.producerSleepIntervalInMillisecond)
				pc.producerSleepIntervalInMillisecond /= 2
				pc.producerSleepIntervalMu.Unlock()
			}
		}
		{
			//TODO: consumer adaptive control
		}
		time.Sleep(1 * time.Second) //TODO: make this configurable
	}

}

func (pc *ProducerConsumer) GracefulShutdown() {
	pc.closing = true
	<-pc.producerStopped
	<-pc.consumerStopped
	close(pc.producerStopped)
	close(pc.consumerStopped)
	fmt.Println("graceful shutdown success")
}
