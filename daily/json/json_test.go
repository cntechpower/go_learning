package json

import (
	"encoding/json"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

type IA interface {
	a() string
}

type A struct {
	S string
}

func (a *A) a() string {
	return a.S
}

var testStr string

func TestJsonInterface(t *testing.T) {
	a := &A{S: "origin a"}
	bs, err := json.Marshal(a)
	assert.Equal(t, nil, err)
	testStr = string(bs)
	t.Logf("testStr is: %v\n", testStr)
	var iA IA
	a1 := &A{}
	iA = a1
	rv := reflect.ValueOf(iA)
	t.Logf("iA kind is: %v\n", rv.Kind())
	err = json.Unmarshal([]byte(testStr), iA)
	assert.Equal(t, nil, err)
	t.Logf("iA is: %v\n", iA)
}
