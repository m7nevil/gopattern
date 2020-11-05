package main

import (
	"github.com/m7nevil/gopattern/behavior"
	"github.com/m7nevil/gopattern/structor"
	"log"
)

type Handler1 struct {
	behavior.BaseHandler
}
func (h *Handler1) Handle(request interface{}) {
	log.Println("Handler1 is Handling")
	h.BaseHandler.Handle(request)
}

type Handler2 struct {
	behavior.BaseHandler
}
func (h *Handler2) Handle(request interface{}) {
	log.Println("Handler2 is Handling")
	h.BaseHandler.Handle(request)
}
type Handler3 struct {
	behavior.BaseHandler
}
func (h *Handler3) Handle(request interface{}) {
	log.Println("Handler3 is Handling")
	h.BaseHandler.Handle(request)
}

func main()  {
	//testChain()
	testDecorator()
}

func testChain()  {
	h1 := &Handler1{}
	h2 := &Handler2{}
	h3 := &Handler3{}

	h1.SetNext(h2)
	h2.SetNext(h3)

	h1.Handle("脑壳帝你几叼")
}
func testDecorator()  {
	tds := structor.NewTextDataSource("Hello World")
	ed := structor.NewEncryptionDecorator(tds)
	ed.WriteData(tds.ReadData())

	pd := structor.NewPrefixDecorator(tds)
	pd.WriteData(ed.ReadData())

	log.Println(pd.ReadData())
}