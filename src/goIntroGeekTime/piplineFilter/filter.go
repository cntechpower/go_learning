package piplineFilter

import "github.com/pkg/errors"

var wrongTypeError = errors.New("wrong type of data")

type Request interface {
}

type Response interface {
}

type Filter interface {
	Process(data Request) (Response, error)
}
