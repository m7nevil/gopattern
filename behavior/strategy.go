package behavior

import "log"

type Strategy interface {
	Execute(data interface{})
}

type Context struct {
	strategy Strategy
}
func (c *Context) SetStrategy(s Strategy)  {
	c.strategy = s
}
func (c *Context) Do(data interface{}) {
	c.strategy.Execute(data)
}

type ConStrategy struct {
	
}

func (cs *ConStrategy) Execute(data interface{})  {
	log.Println("Execute")
}

