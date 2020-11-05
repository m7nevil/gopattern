package behavior

import (
	"testing"
)

type Handler1 struct {
	BaseHandler
}
//func (h *Handler1) Handle(request interface{}) {
//	log.Println("Handler1 is Handling")
//}

type Handler2 struct {
	BaseHandler
}
//func (h *Handler2) Handle(request interface{}) {
//	log.Println("Handler2 is Handling")
//}
type Handler3 struct {
	BaseHandler
}
//func (h *Handler3) Handle(request interface{}) {
//	log.Println("Handler3 is Handling")
//}

func testChain(t *testing.T)  {

	h1 := &Handler1{}
	h2 := &Handler2{}
	h3 := &Handler3{}

	h1.SetNext(h2)
	h2.SetNext(h3)

	h1.Handle("脑壳帝你几叼")
}