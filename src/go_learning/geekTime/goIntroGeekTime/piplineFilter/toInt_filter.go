package piplineFilter

import "strconv"

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (f *ToIntFilter) Process(data Request) (Response, error) {
	dataParts, ok := data.([]string)
	if !ok {
		return nil, wrongTypeError
	}
	ret := make([]int, 0)
	for _, part := range dataParts {
		valueInt, err := strconv.Atoi(part)
		if err != nil {
			return nil, wrongTypeError
		}
		ret = append(ret, valueInt)
	}
	return ret, nil
}
