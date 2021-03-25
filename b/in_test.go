package b

import (
	"encoding/json"
	"testing"
)

type a1 struct {
	Name1 string
	Age1  int64
}

type a2 struct {
	Name2 string
	Age2  int64
	City  string
}

type D interface{}

type Res struct {
	D     `json:"data"`
	Votes int64
}

func TestIn(t *testing.T) {
	a11 := &a1{
		Name1: "dujinyang",
		Age1:  10,
	}
	a12 := &a2{
		Name2: "dujinyang",
		Age2:  10,
		City:  "shanghai",
	}
	res1 := &Res{
		D:     a11,
		Votes: 5,
	}
	res2 := &Res{
		D:     a12,
		Votes: 10,
	}
	//
	//valueOf := reflect.ValueOf(res1)
	//typeOf := valueOf.Type()
	//var structFields []reflect.StructField
	//for i := 0; i < valueOf.NumField(); i++ {
	//	fTyp := typeOf.Field(i)
	//	structFields = append(structFields, fTyp)
	//}
	//structFields = append(structFields, reflect.StructField{
	//	Name: "Votes",
	//	Type: reflect.TypeOf(int64(0)),
	//	Tag:  `json:"votes"`,
	//})
	//reflect.StructOf(structFields)
	bs1, _ := json.Marshal(res1)
	bs2, _ := json.Marshal(res2)
	t.Logf("%v\n", string(bs1))
	t.Logf("%v\n", string(bs2))
}
