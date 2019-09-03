package piplineFilter

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (f *SumFilter) Process(data Request) (Response, error) {
	dataIntArray, ok := data.([]int)
	if !ok {
		return nil, wrongTypeError
	}
	sum := 0
	for _, v := range dataIntArray {
		sum += v
	}
	return sum, nil
}
