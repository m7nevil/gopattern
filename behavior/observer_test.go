package behavior

import "log"

type Editor struct {
	Em EventManager
	data interface{}
}

func (e Editor) setData(data interface{}) {
	e.data = data
	e.Em.notify("setData", data)
}

func NewEditor() *Editor{
	return &Editor{Em: EventManager{}}
}

type LogListener struct {}
func (l LogListener) Update(data interface{}) {
	log.Println("write log:", data)
}

type EmailListener struct {}
func (e EmailListener) Update(data interface{})  {
	log.Println("Send Email:", data)
}
