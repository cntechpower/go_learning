package t2pc

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2Pc(t *testing.T) {
	c := &Coordinator{}
	ps := make([]*Participant, 0)
	for i := 0; i <= 5; i++ {
		ps = append(ps, &Participant{})
	}
	for _, p := range ps {
		tmp := p
		c.AddParticipant(tmp)
	}
	for i := 0; i <= 10000; i++ {
		c.Do(rand.Int() % 10)
	}
	for i := 1; i < len(ps); i++ {
		assert.Equal(t, ps[i-1], ps[i])
	}
	t.Logf("Total Commits: %v", c.Commits)
	t.Logf("Total Rollbacks: %v", c.Rollbacks)
}
