// http middleware 实现 
type Context struct {
	handlers []HandlerFunc
	index    int
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
}

// HandlerFunc
func A(c *Context) {
	part1
	c.Next()
	part2
}
func B(c *Context) {
	part3
	c.Next()
	part4
}

// chain 洋葱模型: part1 -> part3 -> part4 -> part2