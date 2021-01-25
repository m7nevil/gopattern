package ioc

type SuperMan struct {
	ability Ability
}

func NewSuperMan(ability Ability) *SuperMan{
	return &SuperMan{ability: ability}
}
func (sm *SuperMan) UsePower()  {
	sm.ability.Activate()
}
