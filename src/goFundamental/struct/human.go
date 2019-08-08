package main

type human struct {
	Sex int
}

func (h *human) setSex(i int) {
	h.Sex = i
}
