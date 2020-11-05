package structor

import "log"

type Component interface {
	Excecute()
}

type BaseDecorator struct {
	wrappee Component
}

func (bd *BaseDecorator) BaseDecorator(c Component) *BaseDecorator {
	return &BaseDecorator{wrappee: c}
}

func (bd *BaseDecorator) Execute() {
	log.Println("BaseDecorator Execute!")
	bd.wrappee.Excecute()
}

type ConDecorator struct {
	BaseDecorator
}

func (cd *ConDecorator) extra() {
	log.Println("ConDecorator extra!")
}

func (cd *ConDecorator) Execute()  {
	cd.BaseDecorator.Execute()
	cd.extra()
}

