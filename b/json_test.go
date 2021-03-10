package b

import (
	"encoding/json"
	"log"
	"testing"
)

type T struct {
	ScoreID       string `json:"score_id" form:"score_id" validate:"required"`
	ViewGenerated bool   `json:"view_generated,omitempty" form:"-"`
}

func TestJson(t *testing.T) {
	t1 := &T{
		ScoreID:       "abc",
		ViewGenerated: false, //this should ignore in bs
	}
	bs, _ := json.Marshal(t1)
	log.Printf("%v", string(bs))
}
