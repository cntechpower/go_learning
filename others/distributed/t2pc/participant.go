package t2pc

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func shouldFail() bool {
	return (time.Now().UnixNano() % 10) == int64(rand.Int()%10)
}

type IParticipant interface {
	PreCommit(n int) bool
	Rollback(n int) bool
	Commit(n int) bool
}

type Participant struct {
	mu       sync.Mutex
	count    int
	preCount int
}

func (w *Participant) PreCommit(n int) (success bool) {
	if shouldFail() {
		return false
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	w.preCount += n
	return true
}

func (w *Participant) Rollback(n int) (success bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.preCount < n {
		panic(fmt.Errorf("w.preCount rollback to minus 0, current %v, rollback %v", w.preCount, n))
		return false
	}
	w.preCount = w.preCount - n
	return true
}

func (w *Participant) Commit(n int) (success bool) {
	if shouldFail() {
		return false
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.preCount < n {
		return false
	}
	w.count += n
	w.preCount = w.preCount - n
	return true
}
