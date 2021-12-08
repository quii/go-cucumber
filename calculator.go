package go_cucumber

type Calculator struct {
	numbers []int
	result  int
}

func (c *Calculator) EnterNumber(x int) {
	c.numbers = append(c.numbers, x)
}

func (c *Calculator) Add() {
	c.result = c.numbers[0] + c.numbers[1] // naive, idc
}

func (c *Calculator) Multiply() {
	c.result = c.numbers[0] * c.numbers[1] // naive, idc
}

func (c *Calculator) Result() int {
	result := c.result
	c.clear()
	return result
}

func (c *Calculator) clear() {
	c.numbers = []int{}
	c.result = 0
}
