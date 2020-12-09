package t2pc

import (
	"time"
)

type Coordinator struct {
	Participants []IParticipant
	Rollbacks    int64
	Commits      int64
}

func (c *Coordinator) AddParticipant(p IParticipant) {
	c.Participants = append(c.Participants, p)
}

func (c *Coordinator) Do(n int) {
	//pre commit
	preCommitDone := make([]IParticipant, 0, len(c.Participants))
	rollback := func() {
		for _, p := range preCommitDone {
			p.Rollback(n)
		}
	}
	for _, p := range c.Participants {
		tmpP := p
		if !tmpP.PreCommit(n) {
			rollback()
			c.Rollbacks++
			return
		}
		preCommitDone = append(preCommitDone, tmpP)
	}

	//commit
	for _, p := range c.Participants {
		tmpP := p
		for { //we must commit success here
			if tmpP.Commit(n) {
				break
			}
			time.Sleep(time.Millisecond)
		}
	}
	c.Commits++
	return
}
