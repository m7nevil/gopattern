package behavior

import "log"

type Handler interface {
	SetNext(next Handler)
	Handle(request interface{})
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler)  {
	h.next = next
}

func (h *BaseHandler) Handle(request interface{})  {
	log.Println("BaseHandler is handling:", request)
	if h.next != nil {
		h.next.Handle(request)
	}
}
