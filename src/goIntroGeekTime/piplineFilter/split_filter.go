package piplineFilter

import "strings"

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{
		delimiter: delimiter,
	}
}
func (f *SplitFilter) Process(data Request) (Response, error) {
	dataString, ok := data.(string)
	if !ok {
		return nil, wrongTypeError
	}
	return strings.Split(dataString, f.delimiter), nil
}
