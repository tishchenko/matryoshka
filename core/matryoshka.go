package core

import "matryoshka/api"

type Matryoshka struct {
	rest *api.RestApiService
}

func NewMatryoshka() *Matryoshka {
	m := &Matryoshka{}

	m.rest = api.NewRestApiService()

	return m
}

func (c *Matryoshka) Run() {
	c.rest.Run()
}