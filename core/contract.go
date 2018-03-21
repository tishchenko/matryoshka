package core

type Contract struct {
}

func NewContract() *Contract {
	c := &Contract{}

	return c
}

func (c *Contract) Validate() bool {

	return true
}
