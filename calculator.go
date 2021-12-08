package go_cucumber

type Calculator struct {
	numbers []int
	result  int
}

func (c *Calculator) EnterNumber(x int) {
	c.numbers = append(c.numbers, x)
}

func (c *Calculator) Add() {
	c.result = 0
	for _, n := range c.numbers {
		c.result += n
	}
}

func (c *Calculator) Result() int {
	return c.result
}
