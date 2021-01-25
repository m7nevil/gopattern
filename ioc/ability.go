package ioc

import "log"

type Ability interface {
	Activate()
}

type XPower struct {
	
}
func (xp *XPower) Activate()  {
	log.Println("XPower Activated!")
}

type SuperBomb struct {

}
func (xp *SuperBomb) Activate()  {
	log.Println("SuperBomb Activated!")
}

type HolyShit struct {

}
func (xp *HolyShit) Activate()  {
	log.Println("HolyShit Activated!")
}