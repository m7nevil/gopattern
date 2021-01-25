package main

import (
	"github.com/m7nevil/gopattern/behavior"
	"github.com/m7nevil/gopattern/ioc"
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
	//testDecorator()
	testIoc()
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

func testIoc() {
	container := &ioc.Container{}
	container.Bind("superman", func(c *ioc.Container, abName string) *ioc.SuperMan {
		ab, _ := container.Make(abName).(ioc.Ability)
		return ioc.NewSuperMan(ab)
	})

	container.Bind("xpower", func(c *ioc.Container) ioc.Ability {
		return &ioc.XPower{}
	})
	container.Bind("holyshit", func(c *ioc.Container) ioc.Ability {
		return &ioc.HolyShit{}
	})
	container.Bind("superbomb", &ioc.SuperBomb{})

	sp1, _ := container.Make("superman", "xpower").(*ioc.SuperMan)
	sp2, _ := container.Make("superman", "holyshit").(*ioc.SuperMan)
	sp3, _ := container.Make("superman", "superbomb").(*ioc.SuperMan)
	//
	sp1.UsePower()
	sp2.UsePower()
	sp3.UsePower()
}