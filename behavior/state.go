package behavior


type State interface {
	DoThis()
	DoThat()
}

type ConcreteState struct {
	context StateContext
}

func (c ConcreteState) DoThis() {

}
func (c ConcreteState) DoThat() {

}

type StateContext struct {
	state State
}
func NewContext(initState State) *StateContext{
	return &StateContext{
		state: initState,
	}
}

func (c StateContext) changeState(state State) {

}
func (c StateContext) DoThis() {
	c.state.DoThis()
}
func (c StateContext) DoThat() {
	c.state.DoThat()
}